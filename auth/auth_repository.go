package auth

import (
	"database/sql"

	"github.com/pkg/errors"
)

type AuthRepository struct {
	Db *sql.DB
}

type User struct {
	User_id    string `json:"user_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Created_on string `json:"created_on"`
}

type Authorization interface {
	CreateUser(user User) (int, error)
	GetUser() (*User, error)
}

//
func NewAuthRepository(Db *sql.DB) *AuthRepository {
	return &AuthRepository{Db: Db}
}

func (ar *AuthRepository) CreateUser(user User) (int, error) {

	result, err := ar.Db.Exec("insert into users (username,password, created_on) values ($1,$2,$3)", user.Username, user.Password, user.Created_on)
	if err != nil {
		return 0, errors.Wrap(err, "User not created")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "user_id not valid!")
	}
	return int(id), nil
}

func (ar *AuthRepository) GetUser(username string) (*User, error) {

	var user User
	rows := ar.Db.QueryRow("select user_id, username, password from users where username = $1", username)

	err := rows.Scan(&user.User_id, &user.Username, &user.Password)
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "user_id not valid!")
	}

	if user.User_id == "" {
		return nil, errors.New("нет такого пользователя")
	}

	return &user, err
}
