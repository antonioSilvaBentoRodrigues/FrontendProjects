package main

import (
	"net/http"

	"github.com/antonioSilvaBentoRodrigues/Login-GO/handlers"
)

var port string = ":8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", handlers.SignUp)
	mux.HandleFunc("/login", handlers.Login)

	//Serve style sheet
	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	http.ListenAndServe(port, mux)
}
