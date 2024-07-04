package main

import (
	"fmt"
	"net/http"

	ret "go-weather-app/retriever"
)

func handleData(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	data := ret.RetrieveData(key)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func main() {
	http.HandleFunc("/weather", handleData)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error")
		return
	}
}
