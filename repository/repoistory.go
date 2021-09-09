package repository

import (
	"database/sql"
	"log"
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
	Data                string `json:"data"`
}

func (r *UserReqResRepository) FindByRequest(req string) (*[]UserReqRes, error) {
	rows, err := r.Db.Query("SELECT * FROM usersreqres where request = $1", req)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userReqRes []UserReqRes
	for rows.Next() {
		var us UserReqRes

		err := rows.Scan(&us.Data_id, &us.Request, &us.City, &us.Latitude, &us.Longitude,
			&us.Temperature, &us.Weatherdescriptions, &us.Humidity, &us.Data)
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
			&us.Temperature, &us.Weatherdescriptions, &us.Humidity, &us.Data)
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

// что должен возвращать метод create
func (r *UserReqResRepository) CreateUsersReqRes(request, city, latitude, longitude, temperature, weatherdescriptions, humidity, data string) (string, error) {
	sqlStatement := `INSERT INTO usersreqres (request, city, latitude, longitude, temperature, weatherdescriptions, humidity, data) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err := r.Db.Exec(sqlStatement, request, city, latitude, longitude, temperature, weatherdescriptions, humidity, data)
	if err != nil {
		log.Fatal(err)
	}
	return "ok", err
}

//dependency injection]
func NewReqResRepository(Db *sql.DB) *UserReqResRepository {

	return &UserReqResRepository{Db: Db}
}
