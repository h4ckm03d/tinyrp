package server

import (
	"encoding/json"
	"net/http"

	"github.com/h4ckm03d/tinyrp/internal/configs"
)

// ping returns a "pong" message
func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// ping returns a "pong" message
func discoveryService(config *configs.Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		discovery := map[string]bool{}
		var client = &http.Client{}
		for _, resource := range config.Resources {
			discovery[resource.Name] = false
			request, err := http.NewRequest("GET", resource.Destination_URL+resource.Endpoint, nil)
			if err != nil {
				continue
			}

			_, err = client.Do(request)
			if err != nil {
				continue
			}

			discovery[resource.Name] = true
		}

		json.NewEncoder(w).Encode(discovery)
	}
}
