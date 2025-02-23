package exporter

import (
	"fmt"
	"iostat_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// تعریف متریک‌ها
var (
	readOps = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "disk_read_ops",
		Help: "Read operations per second",
	}, []string{"device"})

	writeOps = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "disk_write_ops",
		Help: "Write operations per second",
	}, []string{"device"})

	util = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "disk_utilization",
		Help: "Disk utilization percentage",
	}, []string{"device"})
)

func CollectMetrics() {
	stats, err := collector.RunIostat()
	if err != nil {
		fmt.Println("Error running iostat:", err)
		return
	}

	for _, stat := range stats {
		readOps.WithLabelValues(stat.Device).Set(stat.ReadOps)
		writeOps.WithLabelValues(stat.Device).Set(stat.WriteOps)
		util.WithLabelValues(stat.Device).Set(stat.Util)
	}
}
