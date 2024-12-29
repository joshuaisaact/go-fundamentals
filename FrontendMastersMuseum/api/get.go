package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"frontendmasters.com/go/museum/data"
)

// Send a JSON of all the data
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// api/exhibitions?id=34
	id := r.URL.Query()["id"]
	if id != nil { // We will try to serve one
		index, err := strconv.Atoi(id[0])
		if err == nil && index < len(data.GetAll()) {
			// Return the one valid element
			json.NewEncoder(w).Encode(data.GetAll()[index])
		} else {
			http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
