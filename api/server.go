package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matkinhig/go-blogs/api/router"
)

func Run() {
	fmt.Println(" \n\t Listening [::]:3009 \n")
	r := router.New()
	log.Fatal(http.ListenAndServe(":3009", r))
}
