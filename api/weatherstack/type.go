package weatherstack

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
