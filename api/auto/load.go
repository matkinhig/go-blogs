package auto

import (
	"log"

	"github.com/matkinhig/go-blogs/api/database"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// err = db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = db.Debug().AutoMigrate(&models.User{}, models.Post{}).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for i, _ := range users {
	// 	err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	posts[i].AuthorID = users[i].ID
	// 	err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	console.Pretty(posts[i])
	// }

	//create table product
	// err = db.Debug().AutoMigrate(&models.Product{}, &models.Category{}).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//create table categories
	// err = db.Model(&models.Product{}).AddForeignKey("category_id", "categories(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
