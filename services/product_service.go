package services

import (
	"github.com/ReginaSafarina/gin-firebase-backend/models"
	"github.com/ReginaSafarina/gin-firebase-backend/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{productRepo: repositories.NewProductRepository()}
}