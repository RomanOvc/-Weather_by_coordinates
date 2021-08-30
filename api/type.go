package api

type FullRes struct {
	Region               string   `json:"city"`
	Temperature          int      `json:"temperature"`
	Weather_Descriptions []string `json:"weather_description"`
	Humidity             int      `json:"humidity"`
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
}
