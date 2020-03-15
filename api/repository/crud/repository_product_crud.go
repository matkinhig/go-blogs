package crud

import (
	"github.com/jinzhu/gorm"
	"github.com/matkinhig/go-blogs/api/channels"
	"github.com/matkinhig/go-blogs/api/models"
)

type repositoryProductsCRUD struct {
	db *gorm.DB
}

func NewRepositoryProductsCRUD(db *gorm.DB) *repositoryProductsCRUD {
	return &repositoryProductsCRUD{db}
}

func (r *repositoryProductsCRUD) Save(prod *models.Product) (*models.Product, error) {
	var err error
	db := r.db.Debug()
	defer r.db.Close()
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = db.Model(&models.Product{}).Create(&prod).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return prod, nil
	}
	return prod, db.Commit().Error
}
