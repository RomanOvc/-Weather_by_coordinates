package api

import (
	"WeatherByCoordinates/api/mapbox"
	"WeatherByCoordinates/api/weatherstack"
	"WeatherByCoordinates/repository"
	"fmt"
	"log"
)

type UseCase struct {
	Repo repository.Repository // postgresql
	// grpc
	// redis
}

func (uc *UseCase) FullResult(city string) (*FullRes, error) {
	uc.Repo.GetData()
	res1, err := mapbox.Geocode(city)
	res2, err := weatherstack.Forecast(fmt.Sprintf("%v", res1.Latitude), fmt.Sprintf("%v", res1.Longitude))
	fullData := FullRes{
		Region:               res2.Region,
		Temperature:          res2.Temperature,
		Weather_Descriptions: res2.Weather_Descriptions,
		Humidity:             res2.Humidity,
		Latitude:             res1.Latitude,
		Longitude:            res1.Longitude,
	}

	if err != nil {
		log.Print(err)
	}
	return &fullData, err
}
