package weather

import (
	"fmt"
	"net/http"
)

func RetrieveWeather(lat, lon float64) (*http.Request, error) {
	address := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon%f", lat, lon)
	data, err := http.NewRequest(http.MethodGet, address, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &http.Request{}, err
	}
}
