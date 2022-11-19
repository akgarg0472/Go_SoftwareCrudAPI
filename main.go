package main

import (
	"SoftwareCrudAPI/db"
	softwareRouter "SoftwareCrudAPI/router"
	"SoftwareCrudAPI/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	db.InitDbConnection()
}

func main() {
	loadDotEnv()

	router := loadRouters()

	startServer(router)
}

// function to start web server
func startServer(router *mux.Router) {
	port := utils.GetEnvVariable("PORT", "8080")

	err := http.ListenAndServe("localhost:"+port, router)

	if err != nil {
		panic(err)
	}

	log.Printf("Server is listening on PORT %v", port)
}

// function to load all routers
func loadRouters() *mux.Router {
	return softwareRouter.SoftwareRouter()
}

// function to load .env file
func loadDotEnv() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}
