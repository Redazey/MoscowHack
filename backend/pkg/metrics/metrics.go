package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init(mux *http.ServeMux) {
	mux.Handle("/metrics", promhttp.Handler())
}
