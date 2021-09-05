package repository

import (
	"database/sql"
)

type UserReqResRepository struct {
	Db *sql.DB
}

type UserReqRes struct {
	Data_id             string
	Request             string
	City                string
	Latitude            string
	Longitude           string
	Temperature         string
	Weatherdescriptions string
	Humidity            string
	Data                string
}

func (r *UserReqResRepository) FindByRequest(req string) ([]UserReqRes, error) {
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
	return userReqRes, nil
}
