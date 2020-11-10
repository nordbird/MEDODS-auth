package main

import (
	"./controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signin", controllers.SignIn)
	http.HandleFunc("/refresh", controllers.Refresh)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
