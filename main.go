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

func main() {
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

	rep := repository.NewReqResRepository(db)
	usecase := handlers.NewUseCase(rep)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/weather", usecase.WeatherInfo).Methods("GET")
	// register
	// auth <- jwt
	// header = jwt
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Изучить jwt
// Реализовать аутентификацию с помощью jwt
// унифицировать пользователя по jwt и указать валедльца запроса в бд
// Claims, EXP, key