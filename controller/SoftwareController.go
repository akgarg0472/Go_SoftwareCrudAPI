package controller

import "net/http"

//goland:noinspection ALL
func SoftwareHome(writer http.ResponseWriter, _ *http.Request) {
	writer.Write([]byte("Welcome to Software CRUD API"))
}

//goland:noinspection ALL
func AddSoftware(writer http.ResponseWriter, request *http.Request) {}

//goland:noinspection ALL
func GetSoftware(writer http.ResponseWriter, request *http.Request) {}

//goland:noinspection ALL
func GetAllSoftwares(writer http.ResponseWriter, request *http.Request) {}

//goland:noinspection ALL
func UpdateSoftware(writer http.ResponseWriter, request *http.Request) {}

//goland:noinspection ALL
func DeleteSoftware(writer http.ResponseWriter, request *http.Request) {}
