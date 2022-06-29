package api

import (
	"encoding/json"
	"net/http"

	"github.com/amandinedrebes/orness/internal/core"
)

func readNotes(api *API, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Connection", "close")
	defer req.Body.Close()
	api.SetDefaultHeader(w, req)

	filterBy := req.FormValue("tag")

	results := api.Database.Get(filterBy)

	json.NewEncoder(w).Encode(results)
}

func addNotes(api *API, w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Connection", "close")
	defer req.Body.Close()
	api.SetDefaultHeader(w, req)

	var content core.Message
	err := json.NewDecoder(req.Body).Decode(&content)
	if err != nil {
		api.SendError(w, APIErrorBodyParsing, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if content.Message == "" {
		api.SendError(w, APIErrorBodyParsing, "Invalid input", http.StatusInternalServerError)
		return
	}

	api.Database.Add(content)

	json.NewEncoder(w).Encode(content)
}
