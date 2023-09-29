package routes

import (
	"github.com/akagami-harsh/golangProjects/goBookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("GET")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.Updatebook).Methods("POST")
	router.HandleFunc("/boook/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
