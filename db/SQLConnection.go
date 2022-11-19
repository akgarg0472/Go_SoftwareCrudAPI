package db

import (
	"SoftwareCrudAPI/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var database *sql.DB

func initDatabase() {
	datasource := getDatasource()
	println("Datasource is:", datasource)

	db, err := sql.Open("mysql", datasource)

	if err != nil {
		panic(err)
	}

	database = db

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(15)
}

func getDatasource() string {
	dbUrl := utils.GetEnvVariable("MYSQL_URL", "127.0.0.1:3306")
	username := utils.GetEnvVariable("MYSQL_USERNAME", "root")
	password := utils.GetEnvVariable("MYSQL_PASSWORD", "root")
	dbName := utils.GetEnvVariable("MYSQL_DATABASE_NAME", "sw_crud_db")

	return username + ":" + password + "@tcp(" + dbUrl + ")/" + dbName
}

func ConnectDatabase() (*sql.DB, bool) {
	initDatabase()

	if database != nil {
		return nil, false
	}

	return database, true
}
