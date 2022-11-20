package db

import (
	"SoftwareCrudAPI/model"
	"SoftwareCrudAPI/utils"
	"log"
)

var table = "software"

func initSoftwareDao() {
	_, err := database.Exec("CREATE TABLE IF NOT EXISTS " + table + " (id VARCHAR(255), category VARCHAR(255), title VARCHAR(255), description VARCHAR(255))")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func InsertSoftware(software *model.Software) (bool, error) {
	stmt, err := database.Prepare("INSERT INTO " + table + " (id, category, title, description) VALUES (?, ?, ?, ?)")
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	id := utils.GenerateRandomId()
	category := software.Category
	title := software.Title
	description := software.Description

	insertResult, _ := stmt.Exec(id, category, title, description)
	insertedRows, _ := insertResult.RowsAffected()

	return insertedRows == 1, nil
}

func FindOneById(id string) (model.Software, error) {
	stmt, err := database.Prepare("SELECT id, category, title, description FROM " + table + " WHERE id = ?")
	if err != nil {
		return model.Software{}, err
	}

	defer stmt.Close()

	var software model.Software
	queryRowError := stmt.QueryRow(id).Scan(&software.Id, &software.Category, &software.Title, &software.Description)

	if queryRowError != nil {
		return model.Software{}, err
	}

	return software, nil
}

func FindAll() ([]model.Software, error) {
	stmt, err := database.Prepare("SELECT id, category, title, description FROM " + table)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	queryResult, queryError := stmt.Query()
	if queryError != nil {
		return nil, queryError
	}

	defer queryResult.Close()

	var softwares []model.Software

	for queryResult.Next() {
		var software model.Software
		rowError := queryResult.Scan(&software.Id, &software.Category, &software.Title, &software.Description)

		if rowError != nil {
			log.Println(rowError.Error())
			continue
		}

		softwares = append(softwares, software)
	}

	return softwares, nil
}

func UpdateOne(software model.Software) (bool, error) {
	stmt, err := database.Prepare("UPDATE " + table + " SET category=?, title=?, description=? WHERE id=?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	updateResult, updateError := stmt.Exec(software.Category, software.Title, software.Description, software.Id)

	if updateError != nil {
		return false, updateError
	}

	rowsAffected, _ := updateResult.RowsAffected()

	return rowsAffected == 1, nil
}

func Delete(id string) (bool, error) {
	stmt, err := database.Prepare("DELETE FROM " + table + " WHERE id=?")
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	queryResult, err := stmt.Exec(id)
	if err != nil {
		return false, err
	}

	rowsAffected, _ := queryResult.RowsAffected()

	return rowsAffected == 1, nil
}
