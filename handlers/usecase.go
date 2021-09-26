package handlers

import (
	"WeatherByCoordinates/api/mapbox"
	"WeatherByCoordinates/api/weatherstack"
	"WeatherByCoordinates/repository"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func (repo *UseCase) FullResult(city, user_id string) (*repository.UserReqRes, error) {
	now := time.Now()
	res1, err := mapbox.Geocode(city)
	if err != nil {
		return nil, errors.Wrap(err, "Error request")
	}
	res2, err := weatherstack.Forecast(fmt.Sprintf("%v", res1.Latitude), fmt.Sprintf("%v", res1.Longitude))
	counter, err := repo.Repo.NumberOfRecords()
	if err != nil {
		return nil, errors.Wrap(err, "записей нет")
	}
	intVar, _ := strconv.Atoi(user_id)
	fullData := repository.UserReqRes{
		Data_id:             fmt.Sprint(counter + 1),
		Request:             strings.ToLower(city),
		City:                res2.Region,
		Latitude:            fmt.Sprint(res1.Latitude),
		Longitude:           fmt.Sprint(res1.Longitude),
		Temperature:         fmt.Sprint(res2.Temperature),
		Weatherdescriptions: fmt.Sprint(res2.Weather_Descriptions),
		Humidity:            fmt.Sprint(res2.Humidity),
		User_id:             intVar,
		Data:                fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day()),
	}

	if err != nil {
		return nil, errors.Wrap(err, "Error request")
	}

	return &fullData, err
}

func (repo *UseCase) AddData(structData *repository.UserReqRes) (string, error) {
	err := repo.Repo.CreateUsersReqRes(
		structData.Request,
		structData.City,
		structData.Latitude,
		structData.Longitude,
		structData.Temperature,
		structData.Weatherdescriptions,
		structData.Humidity,
		structData.Data,
		structData.User_id,
	)
	if err != nil {
		return "", errors.Wrap(err, "error datas")
		// return "", err
	}
	return "ok", err

}
