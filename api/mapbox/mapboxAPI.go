package mapbox

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// вводим название города, получаем координаты и передаём их в weatherstack
// https://api.mapbox.com/geocoding/v5/mapbox.places/London.json?access_token='access_token'g&limit=1
var urlMapboxApi = "https://api.mapbox.com/geocoding/v5/mapbox.places/"
var access_token = "pk.eyJ1Ijoicm9tYW5vdmM5NyIsImEiOiJja3Nsc3JkdjAwbGJxMm9wZnZucjN0a2w4In0.jc9pottsKnKB7y9CKltErg"

func Geocode(city string) (*MapboxResp, error) {
	var (
		err error
	)
	req, err := http.NewRequest(http.MethodGet, urlMapboxApi+city+".json?", nil)

	q := req.URL.Query()
	q.Add("access_token", access_token)
	q.Add("limit", "1")
	req.URL.RawQuery = q.Encode()
	res, err := http.Get(req.URL.String())
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Error request")
	}
	var data *Mapbox

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, errors.Wrap(err, "ERRor request")
	}

	result := MapboxResp{Latitude: data.Features[0].Center[1], Longitude: data.Features[0].Center[0]}
	return &result, nil
}
