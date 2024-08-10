package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseBody reads the body from an HTTP request and unmarshals it into the provided struct.
func ParseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, x); err != nil {
		return err
	}

	return nil
}
