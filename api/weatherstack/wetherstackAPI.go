package weatherstack

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Current struct {
	Temperature          int
	Weather_Descriptions []string
	Humidity             int
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
	Temperature          int      `json:"temperature"`
	Weather_Descriptions []string `json:"weather_description"`
	Humidity             int      `json:"humidity"`
}

// http://api.weatherstack.com/current?access_key='my_access_key'&query=latitude,longitude
var urlWeatherStack = "http://api.weatherstack.com/current"
var access_key = "a685cabaa481f94d6d324191b608ee6a"

func Forecast(latitude, longitude string) (*ForecastResp, error) {
	var (
		err error
	)
	req, err := http.NewRequest(http.MethodGet, urlWeatherStack, nil)
	q := req.URL.Query()
	q.Add("access_key", access_key)
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

	result := ForecastResp{Region: data.Location.Region, Temperature: data.Current.Temperature, Humidity: data.Current.Humidity, Weather_Descriptions: data.Current.Weather_Descriptions}
	return &result, nil
}
