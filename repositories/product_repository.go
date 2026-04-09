package repositories

import (
	"github.com/ReginaSafarina/gin-firebase-backend/config"
	"github.com/ReginaSafarina/gin-firebase-backend/models"
)

type ProductRepository struct{}
func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}