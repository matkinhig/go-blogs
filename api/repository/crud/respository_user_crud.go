package crud

import (
	"github.com/jinzhu/gorm"
	"github.com/matkinhig/go-blogs/api/channels"
	"github.com/matkinhig/go-blogs/api/models"
)

type repositoryUsersCRUD struct {
	db *gorm.DB
}

func (r *repositoryUsersCRUD) NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUsersCRUD {
	return &repositoryUsersCRUD{db}
}

func (r *repositoryUsersCRUD) Save(u models.User) (modes.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.User{}).Create(&u).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return u, nil
	}
	return models.User{}, nil
}
