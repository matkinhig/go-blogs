package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matkinhig/go-blogs/api/database"
	"github.com/matkinhig/go-blogs/api/models"
	"github.com/matkinhig/go-blogs/api/repository"
	"github.com/matkinhig/go-blogs/api/repository/crud"
	"github.com/matkinhig/go-blogs/api/responses"

)

func test() {
	fmt.Println("oke")
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := crud.NewRepositoryPostsCRUD(db)
	func(postRepo repository.PostRepository) {
		posts, err := postRepo.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, posts)
	}(repo)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	postIn := models.Post{}
	err = json.Unmarshal(body, &postIn)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	postIn.Prepare()
	err = postIn.Validate()
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

	repo := crud.NewRepositoryPostsCRUD(db)

	func(post repository.PostRepository) {
		p, err := post.Save(postIn)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s %s %d", r.Host, r.RequestURI, p.ID))
		responses.JSON(w, http.StatusCreated, &p)
	}(repo)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPostsCRUD(db)
	func(postRepo repository.PostRepository) {
		p, err := postRepo.FindById(uint32(pid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, p)
		return
	}(repo)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post := models.Post{}

	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	post.Prepare()
	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)
	func(postRepo repository.PostRepository) {
		rows, err := postRepo.Update(uint32(pid), post)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, rows)
		return
	}(repo)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postRepo repository.PostRepository) {
		_, err := postRepo.Delete(uint32(pid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", pid))
		responses.JSON(w, http.StatusNoContent, "")
	}(repo)
}
