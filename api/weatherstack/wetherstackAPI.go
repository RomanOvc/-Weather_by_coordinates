package weatherstack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Current struct {
	Temperature         int
	WeatherDescriptions []string
	Humidity            int
}
type Location struct {
	Region string
}

type ForecastR struct {
	Location Location
	Current  Current
}

type ForecastResp struct {
	Region               string   `json:"city"`
	Temperature         int      `json:"temperature"`
	WeatherDescriptions []string `json:"weather_description"`
	Humidity            int      `json:"humidity"`
}

// http://api.weatherstack.com/current?access_key='my_access_key'&query=latitude,longitude
var urlWeatherStack = "http://api.weatherstack.com/current"
var accessKey = "a685cabaa481f94d6d324191b608ee6a"

func Forecast(latitude, longitude string) (*ForecastResp, error) {
	var (
		err error
	)
	req, err := http.NewRequest(http.MethodGet, urlWeatherStack, nil)
	q := req.URL.Query()
	q.Add("access_key", accessKey)
	q.Add("query", latitude+","+longitude)
	req.URL.RawQuery = q.Encode()
	res, err := http.Get(req.URL.String())
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Error request")
	}

	var data *ForecastR

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, errors.Wrap(err, "ERRor request")
	}

	result := ForecastResp{Region: data.Location.Region, Temperature: data.Current.Temperature, Humidity: data.Current.Humidity, WeatherDescriptions: data.Current.WeatherDescriptions}
	return &result, nil
}
