package registry

import (
	"encoding/json"
	"net/http"
)

type ServiceDiscoveryResponse struct {
	Modules string `json:"modules.v1"`
}

func ServiceDiscoveryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := ServiceDiscoveryResponse{
		Modules: "/terraform/modules/v1/",
	}
	body, _ := json.Marshal(resp)
	_, err := w.Write(body)
	if err != nil {
		w.WriteHeader(500)
	}
}
