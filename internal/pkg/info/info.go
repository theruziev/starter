package info

import (
	"encoding/json"
	"net/http"
)

var Version = "dev"
var Commit = ""

func Information() map[string]any {
	return map[string]any{
		"version": Version,
		"commit":  Commit,
	}
}

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info := Information()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(info); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
