package repository

import (
	"WeatherByCoordinates/auth"
	"database/sql"

	"github.com/pkg/errors"
)

//nтут нормально
type UserReqResRepository struct {
	Db *sql.DB
}

//struct db table
type UserReqRes struct {
	Data_id             string `json:"data_id"`
	Request             string `json:"request"`
	City                string `json:"city"`
	Latitude            string `json:"latitude"`
	Longitude           string `json:"longitude"`
	Temperature         string `json:"temperateure"`
	Weatherdescriptions string `json:"weather_descriptions"`
	Humidity            string `json:"humidity"`
	User_id             int    `json:"user_id"`
	Data                string `json:"data"`
}
type FindByIdType struct {
	Request             string `json:"request"`
	City                string `json:"city"`
	Latitude            string `json:"latitude"`
	Longitude           string `json:"longitude"`
	Temperature         string `json:"temperateure"`
	Weatherdescriptions string `json:"weather_descriptions"`
	Humidity            string `json:"humidity"`
	User_id             string `json:"user_id"`
	Data                string `json:"data"`
}

func (r *UserReqResRepository) FindByRequest(req string) (*UserReqRes, error) {
	rows := r.Db.QueryRow("SELECT * FROM usersreqres where request = $1", req)
	var us UserReqRes
	err := rows.Scan(&us.Data_id, &us.Request, &us.City, &us.Latitude, &us.Longitude,
		&us.Temperature, &us.Weatherdescriptions, &us.Humidity, &us.User_id, &us.Data)
	if err = rows.Err(); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &us, err

}

func (r *UserReqResRepository) AllIn() (*[]UserReqRes, error) {
	rows, err := r.Db.Query("SELECT * FROM usersreqres")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userReqRes []UserReqRes
	for rows.Next() {
		var us UserReqRes

		err := rows.Scan(&us.Data_id, &us.Request, &us.City, &us.Latitude, &us.Longitude,
			&us.Temperature, &us.Weatherdescriptions, &us.Humidity, &us.User_id, &us.Data)
		if err != nil {
			return nil, err
		}
		userReqRes = append(userReqRes, us)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &userReqRes, nil

}

func (r *UserReqResRepository) CreateUsersReqRes(request, city, latitude, longitude, temperature, weatherdescriptions, humidity, data string, user_id int) error {
	sqlStatement := `INSERT INTO usersreqres (request, city, latitude, longitude, temperature, weatherdescriptions, humidity,user_id, data) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	res, err := r.Db.Exec(sqlStatement, request, city, latitude, longitude, temperature, weatherdescriptions, humidity, user_id, data)
	if err != nil {
		return err
	}
	setedRows, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "sql type error")
	}

	if setedRows < 1 {
		return errors.New("table not update")
	}

	return err
}

func (r *UserReqResRepository) NumberOfRecords() (int, error) {
	var counter int
	err := r.Db.QueryRow("SELECT count(*) FROM usersreqres").Scan(&counter)
	if err != nil {
		return 0, errors.Wrap(err, "записей нет")
	}
	return counter, err
}

func (r *UserReqResRepository) FindByIdUser(id string) (*[]FindByIdType, error) {
	rows, err := r.Db.Query("SELECT request,city,latitude,longitude,temperature, weatherdescriptions,humidity,user_id,data FROM usersreqres where user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var findByIdType []FindByIdType
	for rows.Next() {
		var us FindByIdType

		err := rows.Scan(&us.Request, &us.City, &us.Latitude, &us.Longitude, &us.Temperature, &us.Weatherdescriptions, &us.Humidity, &us.User_id, &us.Data)
		if err != nil {
			return nil, err
		}
		findByIdType = append(findByIdType, us)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &findByIdType, nil
}

func (r *UserReqResRepository) GetUser(username string) (*auth.User, error) {
	var user auth.User
	err := r.Db.QueryRow("select user_id, username, password from users where username = $1", username).Scan(&user.User_id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func NewReqResRepository(Db *sql.DB) *UserReqResRepository {
	return &UserReqResRepository{Db: Db}
}
