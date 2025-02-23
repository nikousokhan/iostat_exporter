package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartServer راه‌اندازی سرور HTTP برای جمع‌آوری متریک‌ها
func StartServer() {
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Exporter running on :9100")
	log.Fatal(http.ListenAndServe(":9100", nil))
}
