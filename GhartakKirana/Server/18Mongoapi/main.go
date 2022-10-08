package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adithyahemanth/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Ready to Debug....check port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
