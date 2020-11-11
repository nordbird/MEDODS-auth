package main

import (
	"./controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signin", controllers.SignIn)
	http.HandleFunc("/refresh", controllers.Refresh)
	http.HandleFunc("/delete-one-refresh", controllers.DeleteOneRefreshToken)
	http.HandleFunc("/delete-all-refresh", controllers.DeleteAllRefreshToken)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}
