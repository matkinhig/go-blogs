package auto

import (
	"log"

	database "github.com/matkinhig/go-blogs/api/db"
)

func Load() {
	db, err := database.Connect
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().Drop
}
