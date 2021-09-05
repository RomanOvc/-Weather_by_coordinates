package main

import (
	"WeatherByCoordinates/repository"
	"fmt"
	"log"
	"net/http"

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

type Env struct {
	users repository.UserReqResRepository
}

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

	env := &Env{
		users: repository.UserReqResRepository{Db: db},
	}

	http.HandleFunc("/weather", env.weatherInfo)
	http.ListenAndServe(":8000", nil)
}

func (repo *Env) weatherInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	city := r.URL.Query().Get("city")
	uss, err := repo.users.FindByRequest(city)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, us := range uss {
		fmt.Fprintf(w, "%s, %s, %s", us.Data_id, us.Request, us.City)

	}
}

// res1, err := api.FullResult("New York")
// if err != nil {
// log.Print(err)
// }

// u, err := json.Marshal(res1)
// if err != nil {
// log.Print(err)
// }
// fmt.Println(string(u))
// }
