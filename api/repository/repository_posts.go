package repository

import (
	"github.com/matkinhig/go-blogs/api/models"
)

type PostRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
	FindById(uint32) (models.Post, error)
	Update(uint32, models.Post) (int64, error)
	Delete(post_id, user_id uint32) (int64, error)
}
