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
	var err error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var user auth.User
	json.NewDecoder(r.Body).Decode(&user)

	defer func() {
		if err != nil {
			w.WriteHeader(400)
			bytes, _ := json.Marshal("not registre: " + err.Error())
			w.Write(bytes)
			return
		}
	}()
	userId, err := authR.AuthR.CreateUser(user)
	if err != nil {
		return ///TODO как-то вернуть сообщение что такой юзер существует
	}
	log.Print(userId)
	bytes, err := json.Marshal(userId)
	if err != nil {
		return
	}
	w.Write(bytes)
}

func (authR *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var (
		u   []byte
		err error
	)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var user auth.User
	json.NewDecoder(r.Body).Decode(&user)
	defer func() {
		if err != nil {
			log.Println(err, "Error request")
			w.WriteHeader(400)
			w.Write(nil)
		} else {
			w.Write(u)
		}
	}()
	token, err := authR.CheckToken(&user)
	u, err = json.Marshal(token)
	if err != nil {
		return
	}

	return
}

// не знаю куда засунуть middleware
func IsAuthorized(next http.HandlerFunc) http.Handler {
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
				w.WriteHeader(400)
				w.Write(nil)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(401)
				w.Write(nil)
				return
			}

		} else {
			w.WriteHeader(400)
			w.Write(nil)
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
