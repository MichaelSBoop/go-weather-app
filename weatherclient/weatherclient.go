package weatherclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"location"
	"weather"
)

type Coords struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

type WeatherData struct {
	Main  Main  `json:"main"`
	Coord Coord `json:"coord"`
	Wind  Wind  `json:"wind"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humididty"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}

func CallForData() {
	var coords Coords
	var weatherData WeatherData
	client := &http.Client{}
	locationResponse, err := client.Do(location.RetrieveLocation("http://ip-api.com/json/"))
	if err != nil {
		fmt.Printf("Bad Location Request %d", http.StatusBadRequest)
		return
	}
	defer locationResponse.Body.Close()
	body, err := io.ReadAll(locationResponse.Body)
	if err != nil {
		fmt.Println("Failed to parse response")
		return
	}
	if err = json.Unmarshal([]byte(body), &location); err != nil {
		fmt.Println("Failed to unmarshal data", err)
		return
	}
	weatherResponnse, err := client.Do(weather.RetrieveWeather(location.Lat, location.Lon))
	if err != nil {
		fmt.Printf("Bad Weather Request %d", http.StatusBadRequest)
		return
	}
	defer data.Body.Close()
	body, err := io.ReadAll(data.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &http.Request{}, err
	}

	if err = json.Unmarshal(body, &weatherData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &http.Request{}, err
	}
	resp, err := json.Marshal(&weatherData)
	if err != nil {
		http.Error(err.Error(), http.StatusBadRequest)
		return &http.Request{}, err
	}
}
