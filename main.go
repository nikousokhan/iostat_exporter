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
	// دریافت مقدار آی‌پی و پورت از متغیرهای محیطی (با مقدار پیش‌فرض)
	host := os.Getenv("EXPORTER_HOST")
	if host == "" {
		host = "0.0.0.0" // مقدار پیش‌فرض
	}

	port := os.Getenv("EXPORTER_PORT")
	if port == "" {
		port = "9100" // مقدار پیش‌فرض
	}

	address := fmt.Sprintf("%s:%s", host, port)

	collector.RegisterMetrics()

	// Goroutine برای جمع‌آوری متریک‌ها
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
