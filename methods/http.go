package methods

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
	// data needs to be sent as JSON
	// encode it as JSON first
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	// now, data is ready as json
	// create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	// request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// create client, make the request
	client := &http.Client{}
	res, err := client.Do(req) // the http request created
	// returns an http response and an error
	if err != nil {
		return User{}, err
	}
	// now have the http response as json data
	defer res.Body.Close() // close response body after
	// function returns

	// decode json data
	var user User
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}
