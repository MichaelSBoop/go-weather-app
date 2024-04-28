package location

import (
	"fmt"
	"net/http"
)

// GetLocation makes an API-call to ip-api to recieve data on users whereabouts
func GetLocation(address string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, address, nil)
	if err != nil {
		fmt.Printf("Bad request %d", http.StatusBadRequest)
		return &http.Request{}
	}

	params := req.URL.Query()
	params.Add("fields", "city,lat,lon")
	req.URL.RawQuery = params.Encode()

	return req
}
