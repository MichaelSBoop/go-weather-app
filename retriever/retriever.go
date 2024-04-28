package retriever

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"location"
	"weather"
)

type Coords struct {
	City string  `json:"city,omitempty"`
	Lat  float64 `json:"latitude"`
	Lon  float64 `json:"longitude"`
}

type WeatherData struct {
	Main   Main   `json:"main"`
	Coords Coords `json:"coordinates"`
	Wind   Wind   `json:"wind"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humididty"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}

func RetrieveData() []byte {
	var coords Coords
	var weatherData WeatherData
	client := &http.Client{}
	locationResponse, err := client.Do(location.GetLocation("http://ip-api.com/json/"))
	if err != nil {
		fmt.Printf("Bad Location Request %d", http.StatusBadRequest)
		return []byte{}
	}
	defer locationResponse.Body.Close()
	locationBody, err := io.ReadAll(locationResponse.Body)
	if err != nil {
		fmt.Println("Failed to parse response")
		return []byte{}
	}
	if err = json.Unmarshal([]byte(locationBody), &coords); err != nil {
		fmt.Println("Failed to unmarshal data", err)
		return []byte{}
	}
	weatherResponnse, err := client.Do(weather.GetWeather(coords.Lat, coords.Lon))
	if err != nil {
		fmt.Printf("Bad Weather Request %d", http.StatusBadRequest)
		return []byte{}
	}
	defer weatherResponnse.Body.Close()
	weatherBody, err := io.ReadAll(weatherResponnse.Body)
	if err != nil {
		fmt.Println("Failed to parse response")
		return []byte{}
	}
	if err = json.Unmarshal(weatherBody, &weatherData); err != nil {
		fmt.Println("Failed to unmarshal data", err)
		return []byte{}
	}
	resp, err := json.Marshal(&weatherData)
	if err != nil {
		fmt.Println("Failed to marshal data", err)
		return []byte{}
	}
	return resp
}
