package repository

import (
	"github.com/matkinhig/go-blogs/api/models"
)

type ProductRepository interface {
	Save(*models.Product) (*models.Product, error)
}
