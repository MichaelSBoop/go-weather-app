package main

import (
	"fmt"
	"net/http"

	"weatherclient"
)

func handleData(w http.ResponseWriter, r *http.Request) {
	data := weatherclient.CallForData()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func main() {
	http.HandleFunc("/weather", handleData)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка запуска сервера")
		return
	}
	fmt.Println("Hello, world!")
}
