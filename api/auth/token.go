package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/matkinhig/go-blogs/api/config"
	"github.com/matkinhig/go-blogs/api/utils/console"
)

func CreateToken(user_id uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //muon doi thoi gian jwt expired thi thay doi o day
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString(config.SECRETKEY)
	fmt.Println(str)
	fmt.Println(err)
	return str, err
}

func ExtractToken(r *http.Request) string {
	key := r.URL.Query()
	token := key.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func TokenValidate(r *http.Request) error {
	param := ExtractToken(r)
	token, err := jwt.Parse(fmt.Sprintf("%v", param), func(token *jwt.Token) (interface{}, error) {
		//ALWAYS TEST THE PARAM WANT TO VALIDATE WHICH WE EXPECT -- LUON LUON KIEM TRA BUOC NAY -- WITH ALG
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method : %v", token.Header["alg"])
		}

		//hmacSampleSelect IS variable type []byte which contain the secret key
		return config.SECRETKEY, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		console.Pretty(claims)
		return nil
	}
	return nil
}

func ExtractTokenID(r *http.Request) (uint32, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signning method: %v", token.Header["alg"])
		}
		return config.SECRETKEY, nil
	})
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}
