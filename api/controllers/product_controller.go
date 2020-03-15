package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matkinhig/go-blogs/api/auth"
	"github.com/matkinhig/go-blogs/api/database"
	"github.com/matkinhig/go-blogs/api/models"
	"github.com/matkinhig/go-blogs/api/repository"
	"github.com/matkinhig/go-blogs/api/repository/crud"
	"github.com/matkinhig/go-blogs/api/responses"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	prodIn := models.Product{}
	if err = json.Unmarshal(body, &prodIn); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = prodIn.Validate(); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	role, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusNetworkAuthenticationRequired, errors.New("Cant create product because you dont have a permisson"))
		return
	}

	if role != 1 {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryProductsCRUD(db)

	func(prod repository.ProductRepository) {
		p, err := prod.Save(&prodIn)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s %s %d", r.Host, r.URL.Path, p.ID))
		responses.JSON(w, http.StatusCreated, &p)
	}(repo)

}
