package db

import (
	"SoftwareCrudAPI/model"
	"SoftwareCrudAPI/utils"
	"log"
)

var table = "software"

//goland:noinspection SqlDialectInspection,SqlNoDataSourceInspection
func initSoftwareDao() {
	_, err := database.Exec("CREATE TABLE IF NOT EXISTS " + table + " (id VARCHAR(255), category VARCHAR(255), title VARCHAR(255), description VARCHAR(255))")
	if err != nil {
		log.Fatal(err.Error())
	}
}

//goland:noinspection ALL
func InsertOne(software *model.Software) bool {
	stmt, err := database.Prepare("INSERT INTO " + table + " (id, category, title, description) VALUES (?, ?, ?, ?)")
	if err != nil {
		return false
	}

	defer stmt.Close()

	id := utils.GenerateRandomId()
	category := software.Category
	title := software.Title
	description := software.Description

	insertResult, _ := stmt.Exec(id, category, title, description)
	insertedRows, _ := insertResult.RowsAffected()

	return insertedRows == 1
}
