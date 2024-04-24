package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}

func handleData(w http.ResponseWriter, r *http.Request) {
	var weatherData WeatherData
	//var buf bytes.Buffer
	key := r.URL.Query().Get("appid")
	//fmt.Println(key)
	city := r.URL.Query().Get("q")
	//fmt.Println(city)
	address := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, key)
	data, err := http.Get(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer data.Body.Close()
	//_, err = buf.ReadFrom(data.Body)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// }
	body, err := io.ReadAll(data.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err = json.Unmarshal(body, &weatherData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	resp, err := json.Marshal(&weatherData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func main() {
	http.HandleFunc("/weather", handleData)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка запуска сервера")
	}
	fmt.Println("Hello, world!")
}
