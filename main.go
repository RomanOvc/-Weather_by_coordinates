package main

import (
	"WeatherByCoordinates/handlers"
	"WeatherByCoordinates/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5433"
	username = "postgres"
	dbname   = "weatherbycoordinates"
	sslmode  = "disable"
	password = "acer5800"
)

func weatherApi() {
	db, err := repository.InitPostgresDB(repository.Config{
		Host:     host,
		Port:     port,
		Username: username,
		DBName:   dbname,
		SSLMode:  sslmode,
		Password: password,
	})

	if err != nil {
		log.Fatal(err)
	}

	reposa := repository.UserReqResRepository{Db: db}
	usecase := handlers.UseCase{Repo: reposa}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/weather", usecase.WeatherInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	weatherApi()
}
