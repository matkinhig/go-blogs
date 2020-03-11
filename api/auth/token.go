package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/matkinhig/go-blogs/api/config"

)

func CreateToken(user_id uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //muon doi thoi gian jwt expired thi thay doi o day
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SECRETKEY)
}
