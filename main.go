package main

import (
	"WeatherByCoordinates/auth"
	"WeatherByCoordinates/handlers"
	"WeatherByCoordinates/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// var signingKey = []byte("key")

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

	repAuth := auth.NewAuthRepository(db)
	auth := handlers.NewAuthHandler(repAuth)
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/weather", handlers.IsAuthorized(usecase.WeatherInfo)).Methods("GET")
	router.Handle("/request_by_id", handlers.IsAuthorized(usecase.ReuestsByUserIdHandler)).Methods("GET")

	router.HandleFunc("/create_user", auth.CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", auth.LoginHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
