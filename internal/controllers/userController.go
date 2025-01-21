package controllers

import (
	"github/Trung78z/web-service-gin/internal/services"
	"github/Trung78z/web-service-gin/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Response(c, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	ctx := c.Request.Context()

	user, err := ctrl.service.CreateUser(ctx, req.Name, req.Email)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, false, "Failed to create user", nil)
		return
	}

	utils.Response(c, http.StatusOK, true, "User created successfully", user)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {

		utils.Response(c, http.StatusBadRequest, false, "Invalid user ID", nil)
		return
	}

	ctx := c.Request.Context()

	user, err := ctrl.service.GetUserByID(ctx, int32(id))
	if err != nil {
		utils.Response(c, http.StatusNotFound, false, "User not found", nil)
		return
	}

	utils.Response(c, http.StatusOK, true, "User found", user)

}

func (ctrl *UserController) ListUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := ctrl.service.ListUsers(ctx)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, false, "Failed to fetch users", nil)
		return
	}

	utils.Response(c, http.StatusOK, true, "List of users", users)
}
