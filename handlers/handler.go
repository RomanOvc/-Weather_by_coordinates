package handlers

import (
	"WeatherByCoordinates/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type UseCase struct {
	Repo *repository.UserReqResRepository //postgres
}

//TODO эндпоинт должен записать в базу id пользователя, который был получен из токена
func (repo *UseCase) WeatherInfo(w http.ResponseWriter, r *http.Request) {
	var (
		u []byte
	)
	tokenStr := r.Header.Get("Token")
	token, err := jwt.Parse(tokenStr, nil)
	if token == nil {
		log.Println(err)
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	item, ok := claims["user_id"]
	var user_id string
	if ok {
		user_id = fmt.Sprint(item)

	}

	city := r.URL.Query().Get("city") //param

	uss, err := repo.Repo.FindByRequest(strings.ToLower(city))
	if err != nil {
		return
	}

	if uss.City == "" {
		fullResult, err := repo.FullResult(city, user_id) // TODO need handle error
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

		if err != nil {
			log.Println(err, "Error request")
			w.WriteHeader(400)
			w.Write(nil)
		} else {
			w.Write(u)
		}
	return
}

func (repo *UseCase) RequestsByUserIdHandler(w http.ResponseWriter, r *http.Request) {
	var (
		u []byte
	)
	w.Header().Set("Content-Type", "application/json")
	tokenStr := r.Header.Get("Token")
	token, err := jwt.Parse(tokenStr, nil)
	if token == nil {
		log.Println(err)
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	item, ok := claims["user_id"]
	var user_id string
	if ok {
		user_id = fmt.Sprint(item)

	}
	result, err := repo.Repo.FindByIdUser(user_id) // TODO need handle error
	if err != nil {
		return
	}

	u, err = json.Marshal(result)
	if err != nil {
		return
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

