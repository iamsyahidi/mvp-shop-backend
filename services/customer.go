package services

import (
	"math"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/utils"
	"mvp-shop-backend/repositories"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type customerService struct {
	customerRepository repositories.CustomerRepositoryInterface
}

type CustomerServiceInterface interface {
	CreateCustomer(customer *models.Customer) (res *models.Response, err error)
	GetCustomerById(id string) (res *models.Response, err error)
	GetCustomers(filter map[string][]string) (res *models.Response, err error)
	UpdateCustomer(customer *models.CustomerUpdate) (res *models.Response, err error)
	DeleteCustomer(customer *models.CustomerUpdate) (res *models.Response, err error)
}

func NewCustomerService(customerRepository repositories.CustomerRepositoryInterface) CustomerServiceInterface {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (cs *customerService) CreateCustomer(customer *models.Customer) (res *models.Response, err error) {

	customer.Email = strings.ToLower(customer.Email)
	exists, err := cs.customerRepository.GetCustomerByEmail(customer.Email)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	if exists.ID != "" {
		return &models.Response{
			Code:    http.StatusFound,
			Message: "Email already exist",
		}, nil
	}

	password, err := utils.HashPassword(customer.Password)
	if err != nil {
		return nil, err
	}

	customer.ID = uuid.New().String()
	customer.Password = password
	customer.Status = models.StatusActive
	err = cs.customerRepository.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusCreated,
		Message: "Customer created successfully",
	}, nil
}

func (cs *customerService) GetCustomerById(id string) (res *models.Response, err error) {

	customer, err := cs.customerRepository.GetCustomerById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.Response{
				Code:    http.StatusNotFound,
				Message: "Customer not exist",
			}, nil
		}
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "Customer get successfully",
		Data:    customer,
	}, nil
}

func (cs *customerService) GetCustomers(filter map[string][]string) (res *models.Response, err error) {
	pagination, search := utils.GeneratePaginationFromRequest(filter)
	customers, count, err := cs.customerRepository.GetCustomers(pagination, search)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &models.Response{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}, nil
	}

	data := models.ListCustomer{
		Page:      pagination.Page,
		Limit:     pagination.Limit,
		Total:     int(count),
		TotalPage: int(math.Ceil(float64(count) / float64(pagination.Limit))),
		Customers: customers,
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "Customer list successfully",
		Data:    data,
	}, nil
}

func (cs *customerService) UpdateCustomer(customer *models.CustomerUpdate) (res *models.Response, err error) {
	err = cs.customerRepository.UpdateCustomer(customer)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "Customer updated successfully",
	}, nil
}

func (cs *customerService) DeleteCustomer(customer *models.CustomerUpdate) (res *models.Response, err error) {
	err = cs.customerRepository.DeleteCustomer(customer)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Code:    http.StatusOK,
		Message: "Customer deleted successfully",
	}, nil
}
