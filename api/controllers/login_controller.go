package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/matkinhig/go-blogs/api/auth"
	"github.com/matkinhig/go-blogs/api/models"
	"github.com/matkinhig/go-blogs/api/responses"

)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusExpectationFailed, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	token, err := auth.SignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}
