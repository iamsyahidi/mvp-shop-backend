package services

import (
	"mvp-shop-backend/middleware"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/utils"
	"mvp-shop-backend/repositories"
	"net/http"
)

type authService struct {
	customerRepository repositories.CustomerRepositoryInterface
}

type AuthServiceInterface interface {
	Login(auth *models.AuthLogin) (res *models.Response, err error)
}

func NewAuthService(customerRepository repositories.CustomerRepositoryInterface) AuthServiceInterface {
	return &authService{
		customerRepository: customerRepository,
	}
}

func (as *authService) Login(auth *models.AuthLogin) (res *models.Response, err error) {
	authCust, err := as.customerRepository.GetCustomerByEmail(auth.Email)
	if err != nil {
		return &models.Response{
			Code:    http.StatusBadRequest,
			Message: "Email not valid",
		}, nil
	}

	err = utils.CheckPassword(auth.Password, authCust.Password)
	if err != nil {
		return &models.Response{
			Code:    http.StatusBadRequest,
			Message: "Password not valid",
		}, nil
	}

	customerClaims := models.CustomerClaims{
		ID:     authCust.ID,
		Name:   authCust.Name,
		Email:  authCust.Email,
		Status: authCust.Status,
	}

	token, err := middleware.GenerateToken(customerClaims)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "Customer logged in successfully",
		Data: models.AuthToken{
			Token: token,
		},
	}, nil
}
