package db

import (
	"SoftwareCrudAPI/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var database *sql.DB

func InitDbConnection() {
	db, err := sql.Open("mysql", getDatasource())

	if err != nil {
		panic(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(15)

	database = db

	initDao()
}

func getDatasource() string {
	dbUrl := utils.GetEnvVariable("MYSQL_URL", "127.0.0.1:3306")
	username := utils.GetEnvVariable("MYSQL_USERNAME", "root")
	password := utils.GetEnvVariable("MYSQL_PASSWORD", "root")
	dbName := utils.GetEnvVariable("MYSQL_DATABASE_NAME", "sw_crud_db")

	return username + ":" + password + "@tcp(" + dbUrl + ")/" + dbName
}

func initDao() {
	initSoftwareDao()
}
