package controller

import (
	"SoftwareCrudAPI/db"
	"SoftwareCrudAPI/model"
	"SoftwareCrudAPI/utils"
	"encoding/json"
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

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, _err := writer.Write([]byte(err.Error()))
		handleError(_err)
		return
	}

	if !software.IsValid() {
		writer.WriteHeader(http.StatusBadRequest)
		_, _err := writer.Write([]byte("Invalid request input"))
		handleError(_err)
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
	id := request.URL.Query().Get("id")

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
	jsonError := json.NewEncoder(writer).Encode(software)
	handleError(jsonError)
}

func GetAllSoftwares(writer http.ResponseWriter, _ *http.Request) {
	softwares, err := db.FindAll()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _err := writer.Write([]byte(err.Error()))
		handleError(_err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	jsonError := json.NewEncoder(writer).Encode(softwares)
	handleError(jsonError)
}

func UpdateSoftware(writer http.ResponseWriter, request *http.Request) {
	var software model.Software

	err := json.NewDecoder(request.Body).Decode(&software)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, _err := writer.Write([]byte("Invalid request input"))
		handleError(_err)
		return
	}

	updated, updateError := db.UpdateOne(software)
	if updateError != nil {
		_, _err := writer.Write([]byte(updateError.Error()))
		handleError(_err)
		return
	}

	if updated {
		writer.WriteHeader(http.StatusOK)
		_, _err := writer.Write([]byte("Software record inserted successfully"))
		handleError(_err)
	} else {
		_, _err := writer.Write([]byte("Something went wrong while inserting new software"))
		handleError(_err)
	}
}

func DeleteSoftware(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")

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
		_, _err := writer.Write([]byte("Deletion failed"))
		handleError(_err)
	}
}

func handleError(err error) {
	log.Println(err.Error())
}
