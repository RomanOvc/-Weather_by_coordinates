package auth

import "database/sql"

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
	GetUser()
}

//
func NewAuthRepository(Db *sql.DB) *AuthRepository {
	return &AuthRepository{Db: Db}
}

func (ar *AuthRepository) CreateUser(user User) (int, error) {
	var id int

	query := ar.Db.QueryRow("insert into users (username,password, created_on) values ($1,$2,$3) RETURNING user_id", user.Username, user.Password, user.Created_on)
	if err := query.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (ar *AuthRepository) GetUser(username string) (*User, error) {
	var user User
	err := ar.Db.QueryRow("select user_id, username, password from users where username = $1", username).Scan(&user.User_id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, err
}
