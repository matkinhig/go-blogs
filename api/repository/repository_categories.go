package repository

import "github.com/matkinhig/go-blogs/api/models"

type CategoryRepository interface {
	Save(*models.Category) (*models.Category, error)
	// FindAll() ([]models.Category, error)
	// FindById(uint32) (models.Category, error)
	// Update(uint32, models.Category) (int64, error)
	// Delete(uint32) (int64, error)
}
