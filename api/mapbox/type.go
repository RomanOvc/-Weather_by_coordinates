package mapbox

type Feature struct {
	Center []float64
}

type Mapbox struct {
	Features []Feature
}

type MapboxResp struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
