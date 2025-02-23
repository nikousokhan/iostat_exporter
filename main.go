package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"iostat_exporter/collector"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	host := os.Getenv("EXPORTER_HOST")
	if host == "" {
		host = "0.0.0.0" 
	}

	port := os.Getenv("EXPORTER_PORT")
	if port == "" {
		port = "9100"
	}

	address := fmt.Sprintf("%s:%s", host, port)

	collector.RegisterMetrics()

	go func() {
		for {
			collector.CollectIostatMetrics()
			time.Sleep(5 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Exporter running on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
