package crud

import "github.com/jinzhu/gorm"

type &repositoryUsersCRUD struct {
	db *gorm.DB
}

func (r *repositoryUsersCRUD) NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUsersCRUD {
	return &repositoryUsersCRUD{db}
}

func (r *repositoryUsersCRUD) Save(u models.User) (modes.User, error)  {
	var err error
	done := make(chan bool)
	go func(ch chan <- bool){
		err = r.d
	}(done)
}