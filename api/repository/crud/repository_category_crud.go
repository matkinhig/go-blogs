package crud

import (
	"github.com/jinzhu/gorm"
	"github.com/matkinhig/go-blogs/api/channels"
	"github.com/matkinhig/go-blogs/api/models"
)

type repositoryCategoriesCRUD struct {
	db *gorm.DB
}

func NewRepositoryCategoriesCRUD(db *gorm.DB) *repositoryCategoriesCRUD {
	return &repositoryCategoriesCRUD{db}
}

func (r *repositoryCategoriesCRUD) Save(cate *models.Category) (*models.Category, error) {
	var err error
	db := r.db.Debug()
	defer r.db.Close()
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = db.Model(&models.Category{}).Create(&cate).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return cate, nil
	}
	return cate, db.Commit().Error
}

func (r *repositoryCategoriesCRUD) FindAll() (*models.Category, error) {
	return nil, nil
}
