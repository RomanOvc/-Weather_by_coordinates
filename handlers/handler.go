package handlers

import (
	"WeatherByCoordinates/repository"
	"encoding/json"
	"log"
	"net/http"
)

type UseCase struct {
	Repo repository.UserReqResRepository //postgres
}

func (repo *UseCase) WeatherInfo(w http.ResponseWriter, r *http.Request) {
	var (
		u   []byte
		err error
	)

	defer func() {
		if err != nil {
			log.Println(err, "Error request")
			w.Write([]byte(err.Error()))
		} else {
			w.Write(u)
		}
	}()

	w.Header().Set("Content-Type", "application/json")
	city := r.URL.Query().Get("city") //param
	uss, err := repo.Repo.FindByRequest(city)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	u, err = json.Marshal(uss)
	if err != nil {
		return
	}

	return
}
