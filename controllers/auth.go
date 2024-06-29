package controllers

import (
	"mvp-shop-backend/middleware"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/logger"
	"mvp-shop-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authService services.AuthServiceInterface
}

type AuthControllerInterface interface {
	Login(c *gin.Context)
}

func NewAuthController(authService services.AuthServiceInterface) AuthControllerInterface {
	return &authController{
		authService: authService,
	}
}

// Login godoc
// @Summary Login a customer
// @Description Login a customer
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body models.AuthLogin true "Auth"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /auth/login [post]
func (ac *authController) Login(c *gin.Context) {
	var authLogin models.AuthLogin
	if err := c.ShouldBindJSON(&authLogin); err != nil {
		middleware.Response(c, authLogin, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	auth := models.AuthLogin{
		Email:    authLogin.Email,
		Password: authLogin.Password,
	}

	response, err := ac.authService.Login(&auth)
	if err != nil {
		logger.Err(err.Error())
		middleware.Response(c, authLogin, models.Response{
			Code:    http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	middleware.Response(c, authLogin, *response)
}
