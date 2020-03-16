package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matkinhig/go-blogs/api/database"
	"github.com/matkinhig/go-blogs/api/models"
	"github.com/matkinhig/go-blogs/api/repository"
	"github.com/matkinhig/go-blogs/api/repository/crud"
	"github.com/matkinhig/go-blogs/api/responses"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	cateIn := models.Category{}
	err = json.Unmarshal(body, &cateIn)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = cateIn.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	repo := crud.NewRepositoryCategoriesCRUD(db)

	func(cate repository.CategoryRepository) {
		c, err := cate.Save(&cateIn)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s %s %d", r.Host, r.URL.Path, c.ID))
		responses.JSON(w, http.StatusCreated, &c)
	}(repo)

}
