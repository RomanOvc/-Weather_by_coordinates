package api

import (
	"WeatherByCoordinates/api/mapbox"
	"WeatherByCoordinates/api/weatherstack"
	"fmt"
	"log"
)

type FullRes struct {
	Region              string   `json:"city"`
	Temperature         int      `json:"temperature"`
	WeatherDescriptions []string `json:"weather_description"`
	Humidity            int      `json:"humidity"`
	Latitude            float64  `json:"latitude"`
	Longitude           float64  `json:"longitude"`
}

func FullResult(city string) (*FullRes, error) {

	res1, err := mapbox.Geocode(city) // FIXME
	res2, err := weatherstack.Forecast(fmt.Sprintf("%v", res1.Latitude), fmt.Sprintf("%v", res1.Longitude))
	fullData := FullRes{
		Region:              res2.Region,
		Temperature:         res2.Temperature,
		WeatherDescriptions: res2.WeatherDescriptions,
		Humidity:            res2.Humidity,
		Latitude:            res1.Latitude,
		Longitude:           res1.Longitude,
	}
	if err != nil {
		log.Print(err) // TODO Wrap
	}

	return &fullData, err
}
