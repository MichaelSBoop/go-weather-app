package location

import (
	"fmt"
	"net/http"
)

// RetrieveLocation makes an API-call to ip-api to recieve data on users whereabouts
func RetrieveLocation(address string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, address, nil)
	if err != nil {
		fmt.Printf("Bad request %d", http.StatusBadRequest)
		return &http.Request{}, err
	}
	params := req.URL.Query()
	params.Add("fields", "lat")
	params.Add("fields", "lon")
	return req, nil
}
