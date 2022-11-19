package router

import (
	"SoftwareCrudAPI/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func SoftwareRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/", controller.SoftwareHome).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/softwares", controller.AddSoftware).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/softwares", controller.GetSoftware).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/softwares/{id}", controller.GetAllSoftwares).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/softwares/{id}", controller.UpdateSoftware).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/softwares/{id}", controller.DeleteSoftware).Methods(http.MethodDelete)

	return router
}
