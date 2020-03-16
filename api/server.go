package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/matkinhig/go-blogs/api/auto"
	"github.com/matkinhig/go-blogs/api/config"
	"github.com/matkinhig/go-blogs/api/router"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf(" \n\t Listening [::]:%d \n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Request", "Location"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r), handlers.CORS(headers, methods, origins)(r))
}
