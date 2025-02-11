package server

import (
	"context"
	"net/http"

	"github.com/kuskoman/logstash-exporter/config"
)

func getHealthCheck(logstashURL string) func(http.ResponseWriter, *http.Request) {
	client := &http.Client{}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), config.HttpTimeout)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, logstashURL, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if resp.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
