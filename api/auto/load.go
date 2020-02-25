package auto

import (
	"log"

	"github.com/matkinhig/go-blogs/api/database"
	"github.com/matkinhig/go-blogs/api/models"
	"github.com/matkinhig/go-blogs/api/utils/console"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.User{}).Error

	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, u := range users {
		err = db.Debug().Model(&models.User{}).Create(&u).Error
		if err != nil {
			log.Fatal(err)
		}
		console.Pretty(u)
	}
}
