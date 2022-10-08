package router

import (
	"github.com/adithyahemanth/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/getallMovies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/createMovie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/updateMovie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/deleteMovie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteAllMovies", controller.DeleteAllMovie).Methods("DELETE")
	return router
}
