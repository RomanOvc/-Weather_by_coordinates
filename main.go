package main

import (
	"WeatherByCoordinates/auth"
	"WeatherByCoordinates/handlers"
	"WeatherByCoordinates/repository"
	"log"
	"net/http"

	"github.com/go-redis/redis"
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

const (
	addrRedis = "localhost:6379"
	passwordR = ""
	dbR       = 0
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

	redisInit := redis.NewClient(&redis.Options{
		Addr:     addrRedis,
		Password: passwordR,
		DB:       dbR,
	})

	redisClient := repository.NewRedisRepository(redisInit)

	rep := repository.NewReqResRepository(db)
	usecase := handlers.NewUseCase(rep)

	repAuth := auth.NewAuthRepository(db)
	auth := handlers.NewAuthHandler(repAuth, redisClient)

	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/weather", auth.IsAuthorized(usecase.WeatherInfo)).Methods("GET")
	router.Handle("/request_by_id", auth.IsAuthorized(usecase.ReuestsByUserIdHandler)).Methods("GET")

	router.HandleFunc("/create_user", auth.CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", auth.LoginHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
