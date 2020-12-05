package libs

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"errors"
)

// GetIPInfoData gets JSON data from ipinfo website
func GetIPInfoData() ([]byte, error) {
	url := "https://ipinfo.io/json"
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		return ioutil.ReadAll(response.Body)
	}

	errorMessage := fmt.Sprintf("HTTP Status Code: %d", response.StatusCode)
	err = errors.New(errorMessage)
	
	return nil, err
}