package weather

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetWeather(lat, lon float64, key string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, "https://api.openweathermap.org/data/2.5/weather", nil)
	if err != nil {
		fmt.Printf("Bad request %d", http.StatusBadRequest)
		return &http.Request{}
	}

	params := req.URL.Query()
	params.Add("lat", strconv.FormatFloat(lat, 'E', -1, 32))
	params.Add("lon", strconv.FormatFloat(lon, 'E', -1, 32))
	params.Add("appid", key)
	req.URL.RawQuery = params.Encode()

	return req
}
