package routes

import (
	"github.com/Subodh22/Bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{ID}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{ID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{ID}", controllers.DeleteBook).Methods("DELETE")

}
