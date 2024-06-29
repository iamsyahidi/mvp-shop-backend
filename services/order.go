package services

import (
	"context"
	"fmt"
	"mvp-shop-backend/models"
	"mvp-shop-backend/pkg/utils"
	"mvp-shop-backend/repositories"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"
)

type orderService struct {
	orderRepository   repositories.OrderRepositoryInterface
	productRepository repositories.ProductRepositoryInterface
}

type OrderServiceInterface interface {
	CreateOrder(order *models.Order, orderDetail *[]models.OrderDetail) (res *models.Response, err error)
}

func NewOrderService(orderRepository repositories.OrderRepositoryInterface, productRepository repositories.ProductRepositoryInterface) OrderServiceInterface {
	return &orderService{
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
}

func (os *orderService) CreateOrder(order *models.Order, orderDetail *[]models.OrderDetail) (res *models.Response, err error) {
	// Simulate concurrent order insertions
	var wg sync.WaitGroup
	var mu sync.Mutex

	errs, _ := errgroup.WithContext(context.Background())

	var arrOrderDetail []models.OrderDetail
	var amountOrder float64
	invoice := utils.GenerateInvoice()

	for _, v := range *orderDetail {
		product, err := os.productRepository.GetProductById(v.ProductID)
		if err != nil {
			return &models.Response{
				Code:    http.StatusNotFound,
				Message: "Product Not Found",
			}, err
		}

		if product.Stock == 0 {
			continue
		}

		amountDetail := v.Qty * product.Price
		detail := models.OrderDetail{
			Invoice:   invoice,
			ProductID: v.ProductID,
			Qty:       v.Qty,
			Price:     product.Price,
			Amount:    amountDetail,
			CreatedBy: v.CreatedBy,
		}
		amountOrder += amountDetail
		detail.Status = models.StatusActive
		arrOrderDetail = append(arrOrderDetail, detail)
	}

	order.Invoice = invoice
	order.Amount = amountOrder
	order.Status = models.StatusActive

	wg.Add(1)
	errs.Go(func() error {
		errGo := os.orderRepository.TransactionOrder(order, &arrOrderDetail, &wg, &mu)
		if errGo != nil {
			return fmt.Errorf("error in go routine, TransactionOrder, %+v", errGo)
		}
		return nil
	})

	// Wait for completion and return the first error (if any)
	if err := errs.Wait(); err != nil {
		return nil, err
	}

	wg.Wait()

	return &models.Response{
		Code:    http.StatusCreated,
		Message: "Order created successfully",
	}, nil
}
