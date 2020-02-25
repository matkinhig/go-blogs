package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/matkinhig/go-blogs/api/database"
	"github.com/matkinhig/go-blogs/api/models"
	"github.com/matkinhig/go-blogs/api/repository"
	"github.com/matkinhig/go-blogs/api/repository/crud"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	u := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUserCRUD(db)
	func(u repository.UserRepository) {
		newUser, err := u.Save(u)
	}(repo)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("An user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
