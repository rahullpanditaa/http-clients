package methods

import (
	"bytes"
	"encoding/json"
	"errors"
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

func updateUser(baseUrl, id, apiKey string, data User) (User, error) {
	fullUrl := baseUrl + "/" + id

	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("PUT", fullUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}

func getUserById(baseUrl, id, apiKey string) (User, error) {
	fullUrl := baseUrl + "/" + id

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return User{}, err
	}
	req.Header.Set("X-API-KEY", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil
}

func deleteUser(baseUrl, id, apiKey string) error {
	fullUrl := baseUrl + "/" + id

	req, err := http.NewRequest("DELETE", fullUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-API-KEY", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return errors.New("unable to delete resource")
	} else {
		return nil
	}
}
