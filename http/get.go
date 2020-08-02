package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code: %v", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
