package errs

import (
	"encoding/json"
	"log"
)

type CommonError struct {
	StatusCode uint     `json:"statusCode"`
	Status     string   `json:"status"`
	Messages   []string `json:"messages"`
}

func (e *CommonError) Error() string {
	jsonData, err := json.Marshal(e)
	if err != nil {
		log.Fatal("failed to parse error struct", err)
	}

	return string(jsonData)
}
