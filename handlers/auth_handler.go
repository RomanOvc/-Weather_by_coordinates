package handlers

import (
	"WeatherByCoordinates/auth"
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type AuthHandler struct {
	AuthR *auth.AuthRepository
}

func NewAuthHandler(AuthR *auth.AuthRepository) *AuthHandler {
	return &AuthHandler{AuthR: AuthR}
}

// ожидает json следующего вида
//  {
//  	"username":"lol123",
//  	"password":"123",
//  	"created_on":"2021-06-06"
//  }
func (authR *AuthHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var user *auth.User
	json.NewDecoder(r.Body).Decode(&user)

	add, err := authR.AuthR.CreateUser(*user)
	if err != nil {
		return //TODO как-то вернуть сообщение что такой юзер существует
	}
	log.Print(add)
}

func (authR *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var (
		u []byte
	)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var user *auth.User
	json.NewDecoder(r.Body).Decode(&user)

	usse, err := authR.CheckLogin(user)
	u, err = json.Marshal(usse)
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

// не знаю куда засунуть middleware
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Connection", "close")
		defer r.Body.Close()

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				w.Header().Add("Content-Type", "application/json")
				return
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
