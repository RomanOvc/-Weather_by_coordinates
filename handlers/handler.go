package handlers

import (
	"WeatherByCoordinates/repository"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type UseCase struct {
	Repo *repository.UserReqResRepository //postgres
}

func (repo *UseCase) WeatherInfo(w http.ResponseWriter, r *http.Request) {
	var (
		u []byte
	)

	w.Header().Set("Content-Type", "application/json")
	city := r.URL.Query().Get("city") //param

	uss, err := repo.Repo.FindByRequest(strings.ToLower(city))
	if err != nil {
		return
	}

	if uss.City == "" {
		fullResult, err := repo.FullResult(city) // TODO need handle error
		if err != nil {
			return
		}
		u, err = json.Marshal(fullResult)
		if err != nil {
			return
		}

		go repo.AddData(fullResult)

	} else {
		u, err = json.Marshal(uss)
		if err != nil {
			return
		}
	}

	defer func() {
		if err != nil {
			log.Println(err, "Error request")
			w.WriteHeader(400)
			w.Write(nil)
		} else {
			w.Write(u)
		}
	}()
	return
}

func NewUseCase(Repo *repository.UserReqResRepository) *UseCase {
	return &UseCase{Repo: Repo}
}
