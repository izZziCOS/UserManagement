package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/izzzicos/UserManagement/api/internal/handler"
	"github.com/izzzicos/UserManagement/api/internal/models"
	"github.com/izzzicos/UserManagement/api/internal/repository"
	"github.com/izzzicos/UserManagement/api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    // Load configuration
    if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

    // Initialize logger
    logger, err := zap.NewProduction()
    if err != nil {
        log.Fatalf("failed to initialize logger: %v", err)
    }
    defer logger.Sync()

    // Initialize database
    db, err := gorm.Open(postgres.Open(buildDSN()), &gorm.Config{})
    if err != nil {
        logger.Fatal("failed to connect to database", zap.Error(err))
    }

    // Migrate the schema
    if err := db.AutoMigrate(&models.User{}); err != nil {
        logger.Fatal("failed to migrate database", zap.Error(err))
    }

    // Initialize layers
    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    // Create Gin router
    router := gin.Default()

    // Register routes
    userHandler.RegisterRoutes(router)

    // Add health check endpoint
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    // Start server
    srv := &http.Server{
        Addr:    ":" + os.Getenv("API_PORT"),
        Handler: router,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Fatal("failed to start server", zap.Error(err))
        }
    }()

    logger.Info("server started", zap.String("port", os.Getenv("API_PORT")))

    // Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    logger.Info("shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        logger.Error("server forced to shutdown", zap.Error(err))
    }

    logger.Info("server exited properly")
}

func buildDSN() string {
    return "host=" + os.Getenv("DB_HOST") +
        " user=" + os.Getenv("DB_USER") +
        " password=" + os.Getenv("DB_PASSWORD") +
        " dbname=" + os.Getenv("DB_NAME") +
        " port=" + os.Getenv("DB_PORT") +
        " sslmode=disable TimeZone=UTC"
}