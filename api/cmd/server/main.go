package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/izzzicos/UserManagement/api/internal/config"
	"github.com/izzzicos/UserManagement/api/internal/handler"
	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/izzzicos/UserManagement/api/internal/repository"
	"github.com/izzzicos/UserManagement/api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDB() (*gorm.DB, error) {
	dsn := config.BuildDSN()
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func setupRouter(db *gorm.DB, notifier service.Notifier) *gin.Engine {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, notifier)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	userHandler.RegisterRoutes(router)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return router
}

func main() {
	// Configure logger to include timestamps and microsecond precision
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	// Load configuration
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Initialize database
	db, err := setupDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Create notifier
	amqpNotifier, err := service.NewAMQPNotifier(
		os.Getenv("AMQP_URL"),
		"user_events",
	)
	if err != nil {
		log.Fatalf("failed to create notifier: %v", err)

	}
	defer func() {
		if err := amqpNotifier.Close(); err != nil {
			log.Printf("failed to close notifier: %v", err)
		}
	}()

	// Start test consumer in debug mode
	if os.Getenv("DEBUG") == "true" {
		log.Println("starting test consumer for debugging")
		amqpNotifier.StartTestConsumer()
	}

	router := setupRouter(db, amqpNotifier)

	// Start server
	srv := &http.Server{
		Addr:    ":" + os.Getenv("API_PORT"),
		Handler: router,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Server starting on port %s (AMQP queue: %s)",
			os.Getenv("API_PORT"),
			"user_events",
		)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server forced to shutdown: %v", err)
	}

	log.Printf("server exited properly")
}
