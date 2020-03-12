package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matkinhig/go-blogs/api/auth"
	"github.com/matkinhig/go-blogs/api/responses"
)

func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("")
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application-json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValidate(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
