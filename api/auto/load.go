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

	err = db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error

	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, models.Post{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	// var user models.User
	// var post models.Post

	// a := db.Model(&user).Association("ID").Find(&post)

	// fmt.Println(a)

	// err = db.Model(&models.User{}).Related(&models.Post{}).Error
	// // err = db.Debug().Model(&models.Post{}).Where("author_id = ?", &posts[i].AuthorID).Take(&posts[i].Author).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal(err)
		}
		posts[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatal(err)
		}

		console.Pretty(posts[i])
	}
}
