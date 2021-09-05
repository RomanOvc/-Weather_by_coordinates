package repository

import (
	"database/sql"
	"fmt"
)

type UserReqResRepository struct {
	Db *sql.DB
}

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
		userReqRes = append(userReqRes, UserReqRes{
			Data_id:             us.Data_id,
			Request:             us.Request,
			City:                us.City,
			Latitude:            us.Latitude,
			Longitude:           us.Longitude,
			Temperature:         us.Temperature,
			Weatherdescriptions: us.Weatherdescriptions,
			Humidity:            us.Humidity,
			Data:                us.Data,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println(userReqRes)
	return &userReqRes, nil
}
