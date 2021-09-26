package handlers

import (
	"WeatherByCoordinates/auth"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
)

//sign key
var mySigningKey = []byte("ROMAN")

type Token struct {
	Token string `json:"token"`
}

func (authR *AuthHandler) CheckToken(u *auth.User) (*Token, error) {
	user, err := authR.AuthR.GetUser(u.Username)
	if err != nil {
		return nil, err
	}
	if user.Username != u.Username || user.Password != u.Password {
		fmt.Println("NOT CORRECT")
		err := "error"
		return nil, errors.Errorf(err)
	}

	valideToken, err := GenerateJWT(user)
	// fmt.Println(valideToken)
	if err != nil {
		fmt.Println(err)
	}

	return &Token{valideToken}, err

}

// генерация токена
func GenerateJWT(u *auth.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = u.User_id
	claims["username"] = u.Username
	claims["password"] = u.Password

	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		errors.Wrapf(err, "err token")
	}
	return tokenString, err
}
