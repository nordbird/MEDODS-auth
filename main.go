package main

import (
	"log"
	"medods-auth/controllers"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", controllers.Hello)
	http.HandleFunc("/signin", controllers.SignIn)
	http.HandleFunc("/refresh", controllers.Refresh)
	http.HandleFunc("/delete-one-refresh", controllers.DeleteOneRefreshToken)
	http.HandleFunc("/delete-all-refresh", controllers.DeleteAllRefreshToken)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
