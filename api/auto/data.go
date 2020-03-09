package auto

import models "github.com/matkinhig/go-blogs/api/models"

var users = []models.User{
	models.User{Nickname: "matkinhig", Email: "matkinhig@gmail.com", Password: "123456"},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title",
		Content: "Hello Word",
	},
}
