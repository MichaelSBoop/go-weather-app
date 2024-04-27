package weather

import (
	"fmt"
	"net/http"
	"strconv"
)

func RetrieveWeather(lat, lon float64) *http.Request {
	req, err := http.NewRequest(http.MethodGet, "https://api.openweathermap.org/data/2.5/weather", nil)
	if err != nil {
		fmt.Printf("Bad request %d", http.StatusBadRequest)
		return &http.Request{}
	}

	params := req.URL.Query()
	params.Add("lat", strconv.FormatFloat(lat, 'E', -1, 32))
	params.Add("lon", strconv.FormatFloat(lon, 'E', -1, 32))
	params.Add("appid", "d02deaa448e08bd229d0f6f0ff0527ed")
	req.URL.RawQuery = params.Encode()

	return req
}
