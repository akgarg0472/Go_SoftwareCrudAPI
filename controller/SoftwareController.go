package controller

import (
	"SoftwareCrudAPI/db"
	"SoftwareCrudAPI/model"
	"SoftwareCrudAPI/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SoftwareHome(writer http.ResponseWriter, _ *http.Request) {
	_, err := writer.Write([]byte("Welcome to Software CRUD API"))
	handleError(err)
}

func AddSoftware(writer http.ResponseWriter, request *http.Request) {
	var software model.Software

	err := json.NewDecoder(request.Body).Decode(&software)

	bodyValidated := isRequestBodyValidated(&software, &writer, &err)
	if !bodyValidated {
		return
	}

	inserted, insertErr := db.InsertSoftware(&software)

	if insertErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _err := writer.Write([]byte(insertErr.Error()))
		handleError(_err)
		return
	}

	if inserted {
		writer.WriteHeader(http.StatusCreated)
		_, _err := writer.Write([]byte("Software record inserted successfully"))
		handleError(_err)
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _err := writer.Write([]byte("Something went wrong while inserting new software"))
		handleError(_err)
	}
}

func GetSoftware(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]

	if !utils.IsStringValid(&id) {
		writer.WriteHeader(http.StatusBadRequest)
		_, _err := writer.Write([]byte("Invalid ID provided"))
		handleError(_err)
		return
	}

	software, err := db.FindOneById(id)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, _err := writer.Write([]byte(err.Error()))
		handleError(_err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	if software.IsValid() {
		jsonError := json.NewEncoder(writer).Encode(software)
		handleError(jsonError)
	} else {
		jsonError := json.NewEncoder(writer).Encode("No record found with id provided")
		handleError(jsonError)
	}
}

func GetAllSoftwares(writer http.ResponseWriter, _ *http.Request) {
	softwares, err := db.FindAll()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _err := writer.Write([]byte(err.Error()))
		handleError(_err)
		return
	}

	if softwares == nil {
		softwares = make([]model.Software, 0)
	}

	writer.Header().Set("Content-Type", "application/json")
	jsonError := json.NewEncoder(writer).Encode(softwares)
	handleError(jsonError)
}

func UpdateSoftware(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	var software model.Software

	err := json.NewDecoder(request.Body).Decode(&software)

	bodyValidated := isRequestBodyValidated(&software, &writer, &err)
	if !bodyValidated {
		return
	}

	updated, updateError := db.UpdateOne(id, software)
	if updateError != nil {
		_, _err := writer.Write([]byte(updateError.Error()))
		handleError(_err)
		return
	}

	if updated {
		writer.WriteHeader(http.StatusOK)
		_, _err := writer.Write([]byte("Software record updated successfully"))
		handleError(_err)
	} else {
		_, _err := writer.Write([]byte("Something went wrong while updating software"))
		handleError(_err)
	}
}

func DeleteSoftware(writer http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]

	if !utils.IsStringValid(&id) {
		writer.WriteHeader(http.StatusBadRequest)
		_, _err := writer.Write([]byte("Invalid ID provided"))
		handleError(_err)
		return
	}

	deleted, err := db.Delete(id)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _err := writer.Write([]byte(err.Error()))
		handleError(_err)
		return
	}

	if deleted {
		_, _err := writer.Write([]byte("Software record deleted successfully"))
		handleError(_err)
	} else {
		_, _err := writer.Write([]byte("Deletion failed, no record found"))
		handleError(_err)
	}
}

// helper method
func isRequestBodyValidated(software *model.Software, writer *http.ResponseWriter, err *error) bool {
	_writer := *writer

	if !software.IsValid() {
		_writer.WriteHeader(http.StatusBadRequest)
		_, _err := _writer.Write([]byte("Invalid request input(s)"))
		handleError(_err)
		return false
	}

	if err != nil && *err != nil {
		switch {
		case (*err).Error() == "EOF":
			_writer.WriteHeader(http.StatusBadRequest)
			_, _err := _writer.Write([]byte("No Request Body to process"))
			handleError(_err)
			return false

		case err != nil:
			_writer.WriteHeader(http.StatusBadRequest)
			_, _err := _writer.Write([]byte((*err).Error()))
			handleError(_err)
			return false
		}
	}

	return true
}

// helper method
func handleError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
