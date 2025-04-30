package handler

import (
	"net/http"
	"strconv"

	"github.com/izzzicos/UserManagement/api/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
    service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
    router.GET("/users", h.GetUsers)
    router.POST("/users", h.CreateUser)
    router.PUT("/users/:id", h.UpdateUser)
    router.DELETE("/users/:id", h.DeleteUser)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
    filter := service.UserFilter{
        Country: c.Query("country"),
    }

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    pagination := service.Pagination{
        Page:  page,
        Limit: limit,
    }

    users, err := h.service.GetUsers(c.Request.Context(), filter, pagination)
    if err != nil {
        errorResponse(c, http.StatusInternalServerError, "failed to get users")
        return
    }

    c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var req service.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        errorResponse(c, http.StatusBadRequest, "invalid request body")
        return
    }

    user, err := h.service.CreateUser(c.Request.Context(), req)
    if err != nil {
        errorResponse(c, http.StatusInternalServerError, "failed to create user")
        return
    }

    c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
    id := c.Param("id")
    if _, err := uuid.Parse(id); err != nil {
        errorResponse(c, http.StatusBadRequest, "invalid user ID")
        return
    }

    var req service.UpdateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        errorResponse(c, http.StatusBadRequest, "invalid request body")
        return
    }

    user, err := h.service.UpdateUser(c.Request.Context(), id, req)
    if err != nil {
        errorResponse(c, http.StatusInternalServerError, "failed to update user")
        return
    }

    c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if _, err := uuid.Parse(id); err != nil {
        errorResponse(c, http.StatusBadRequest, "invalid user ID")
        return
    }

    if err := h.service.DeleteUser(c.Request.Context(), id); err != nil {
        errorResponse(c, http.StatusInternalServerError, "failed to delete user")
        return
    }

    c.Status(http.StatusNoContent)
}

func errorResponse(c *gin.Context, status int, message string) {
    c.JSON(status, gin.H{"error": message})
}