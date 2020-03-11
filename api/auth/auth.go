package auth

import (
	"github.com/matkinhig/go-blogs/api/database"
	"github.com/matkinhig/go-blogs/api/models"
	"github.com/matkinhig/go-blogs/api/security"

)

func SignIn(email, password string) (string, error) {
	db, err := database.Connect()
	if err != nil {
		return "", err
	}
	defer db.Close()

	user := models.User{}
	err = db.Debug().Model(&models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = security.VerifyPassword(user.Password, password)
	if err != nil {
		return "", err
	}
	return CreateToken(user.ID)
}
