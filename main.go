package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/snjeev-kushwaha/mongoDbApi/router"
)

func main() {
	fmt.Println("Welcome to MongoDb Apis")
	fmt.Println("Server is getting started....")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on Port http://localhost:4000 .......")
}
