package methods

import (
	"encoding/json"
	"net/http"
)

func getusers(url string) ([]User, error) {
	resp, err := http.Get(url) // recieve an http json response
	if err != nil {
		return nil, err
	}

	var users []User
	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}
