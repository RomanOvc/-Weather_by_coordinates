package handlers

import (
	"WeatherByCoordinates/api"
	"WeatherByCoordinates/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type UseCase struct {
	Repo *repository.UserReqResRepository //postgres
}

func (repo *UseCase) WeatherInfo(w http.ResponseWriter, r *http.Request) {
	var (
		u   []byte
		err error
	)

	w.Header().Set("Content-Type", "application/json")
	city := r.URL.Query().Get("city") //param

	uss, err := repo.Repo.FindByRequest(strings.ToLower(city))
	if err != nil {
		return
	}

	if len(*uss) == 0 {
		// дернуть api
		// записать в бд
		// вернуть результат
		now := time.Now()

		kek, err := api.FullResult(city)
		lol, err := repo.Repo.CreateUsersReqRes(
			strings.ToLower(city),
			kek.Region,
			fmt.Sprintf("%f", kek.Latitude),
			fmt.Sprintf("%f", kek.Longitude),
			fmt.Sprint(kek.Temperature),
			fmt.Sprint(kek.Weather_Descriptions),
			fmt.Sprint(kek.Humidity),
			fmt.Sprint(now.Format("2006-01-01")))
		if err != nil {
			return
		}
		log.Print(lol)

		uss, err := repo.Repo.FindByRequest(city)
		u, err = json.Marshal(uss)
		if err != nil {
			return
		}

	} else {
		u, err = json.Marshal(uss)
		if err != nil {
			return
		}
	}

	defer func() {
		if err != nil {
			log.Println(err, "Error request")

		} else {
			w.Write(u)
		}
	}()

	return
}

//dependency injections
func NewUseCase(Repo *repository.UserReqResRepository) *UseCase {
	return &UseCase{Repo: Repo}
}
