package services

import (
	"mvp-shop-backend/models"
	"mvp-shop-backend/repositories"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type cartService struct {
	cartRepository repositories.CartRepositoryInterface
}

type CartServiceInterface interface {
	CreateCart(cart *models.Cart) (res *models.Response, err error)
	GetCartByCustomerID(id string) (res *models.Response, err error)
	DeleteCart(cart *models.CartUpdate) (res *models.Response, err error)
}

func NewCartService(cartRepository repositories.CartRepositoryInterface) CartServiceInterface {
	return &cartService{
		cartRepository: cartRepository,
	}
}

func (cs *cartService) CreateCart(cart *models.Cart) (res *models.Response, err error) {
	cart.ID = uuid.New().String()
	cart.Status = models.StatusActive
	err = cs.cartRepository.CreateCart(cart)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusCreated,
		Message: "Cart created successfully",
	}, nil
}

func (cs *cartService) GetCartByCustomerID(id string) (res *models.Response, err error) {

	carts, err := cs.cartRepository.GetCartByCustomerID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.Response{
				Code:    http.StatusNotFound,
				Message: "Cart not exist",
			}, nil
		}
		return nil, err
	}

	var totalAmount float64
	for _, cart := range carts {
		totalAmount += cart.Amount
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "Cart get successfully",
		Data:    models.CartView{Products: carts, TotalAmount: totalAmount},
	}, nil
}

func (cs *cartService) DeleteCart(cart *models.CartUpdate) (res *models.Response, err error) {
	err = cs.cartRepository.DeleteCart(cart)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "Cart deleted successfully",
	}, nil
}
