// [19/02/2024 - Jugal Suthar ]
package main

import (
	"book-app-server/src/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", controller.CreateBook).Methods("POST")
	r.HandleFunc("/books", controller.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controller.GetBookByID).Methods("GET")
	r.HandleFunc("/books/{id}", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controller.DeleteBook).Methods("DELETE")
	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}).Handler(r)
	http.Handle("/", corsHandler)
	port := "5001"
	fmt.Println("Server is running on :", port)
	http.ListenAndServe(":"+port, nil)

}
