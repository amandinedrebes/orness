package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amandinedrebes/orness/internal/core"
)

const (
	//APIErrorBodyParsing parsing body issue
	APIErrorBodyParsing = 1
	//APIErrorDatabase Database connexion issue
	APIErrorDatabase = 2
	//APIErrorInvalidValue invalid parameter
	APIErrorInvalidValue = 3

	//URLNotes
	URLNotes = "/notes"
)

//APIError Message error code
type APIError struct {
	Code    int    `json:"code"` //errorCode
	Message string `json:"message"`
}

//API description of the web API
type API struct {
	APIIP    string
	APIPort  string
	Database *core.Database
}

type networkError struct {
	s string
}

func (e *networkError) Error() string {
	return e.s
}

// NewError raise an error
func NewError(text string) error {
	return &networkError{text}
}

//SetDefaultHeader set http headers
func (api *API) SetDefaultHeader(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, *")
	w.Header().Set("Content-Type", "application/json")
}

//SendError standardization of API Error message
func (api *API) SendError(w http.ResponseWriter, errorCode int, message string, httpStatus int) {
	errCode := APIError{
		Code:    errorCode,
		Message: message,
	}

	inrec, _ := json.MarshalIndent(errCode, "", "  ")
	fmt.Errorf("API Error: %s", errCode.Message)
	http.Error(w, string(inrec), httpStatus)
}
