package repository

import (
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
	Data                string `json:"data"`
}

func (r *UserReqResRepository) FindByRequest(req string) (*UserReqRes, error) {
	rows := r.Db.QueryRow("SELECT * FROM usersreqres where request = $1", req)
	var us UserReqRes
	err := rows.Scan(&us.Data_id, &us.Request, &us.City, &us.Latitude, &us.Longitude,
		&us.Temperature, &us.Weatherdescriptions, &us.Humidity, &us.Data)
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

func (r *UserReqResRepository) CreateUsersReqRes(request, city, latitude, longitude, temperature, weatherdescriptions, humidity, data string) error {
	sqlStatement := `INSERT INTO usersreqres (request, city, latitude, longitude, temperature, weatherdescriptions, humidity, data) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	res, err := r.Db.Exec(sqlStatement, request, city, latitude, longitude, temperature, weatherdescriptions, humidity, data)
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

// метод, который выведит количесвто записей в таблице, что бы записать его в метод FullResultвозращающий json
func (r *UserReqResRepository) NumberOfRecords() (int, error) {
	var counter int
	err := r.Db.QueryRow("SELECT count(*) FROM usersreqres").Scan(&counter)
	if err != nil {
		return 0, errors.Wrap(err, "записей нет")
	}
	return counter, err
}

// NewReqResRepository constructor
func NewReqResRepository(Db *sql.DB) *UserReqResRepository {
	return &UserReqResRepository{Db: Db}
}
