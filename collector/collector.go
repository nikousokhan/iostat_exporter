package collector

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)



// تعریف متریک‌های Prometheus
var (
	iostatReadRequests = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_read_requests", Help: "Read requests per second"},
		[]string{"device"},
	)
	iostatReadKB = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_read_kb_per_sec", Help: "Kilobytes read per second"},
		[]string{"device"},
	)
	iostatWriteRequests = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_write_requests", Help: "Write requests per second"},
		[]string{"device"},
	)
	iostatWriteKB = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_write_kb_per_sec", Help: "Kilobytes written per second"},
		[]string{"device"},
	)
	iostatUtilization = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_utilization", Help: "Disk utilization percentage"},
		[]string{"device"},
	)
	iostatRRQMS = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_rrqm_per_sec", Help: "Read requests merged per second"},
		[]string{"device"},
	)
	iostatPRRQMS = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_percent_rrqm", Help: "Percentage of read requests merged"},
		[]string{"device"},
	)
	iostatRAwait = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_r_await", Help: "Average read request wait time"},
		[]string{"device"},
	)
	iostatRAreqSz = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_rareq_sz", Help: "Average read request size"},
		[]string{"device"},
	)
	iostatWRQMS = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_wrqm_per_sec", Help: "Write requests merged per second"},
		[]string{"device"},
	)
	iostatPWRQMS = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_percent_wrqm", Help: "Percentage of write requests merged"},
		[]string{"device"},
	)
	iostatWAwait = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_w_await", Help: "Average write request wait time"},
		[]string{"device"},
	)
	iostatWAreqSz = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_wareq_sz", Help: "Average write request size"},
		[]string{"device"},
	)
	iostatAvgQueueSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "iostat_avg_queue_size", Help: "Average queue size"},
		[]string{"device"},
	)	
)

// RegisterMetrics ثبت متریک‌ها در Prometheus
func RegisterMetrics() {
	prometheus.MustRegister(iostatReadRequests)
	prometheus.MustRegister(iostatReadKB)
	prometheus.MustRegister(iostatWriteRequests)
	prometheus.MustRegister(iostatWriteKB)
	prometheus.MustRegister(iostatUtilization)
	prometheus.MustRegister(iostatRRQMS)
	prometheus.MustRegister(iostatPRRQMS)
	prometheus.MustRegister(iostatRAwait)
	prometheus.MustRegister(iostatRAreqSz)
	prometheus.MustRegister(iostatWRQMS)
	prometheus.MustRegister(iostatPWRQMS)
	prometheus.MustRegister(iostatWAwait)
	prometheus.MustRegister(iostatWAreqSz)
	prometheus.MustRegister(iostatAvgQueueSize)
}

// CollectIostatMetrics جمع‌آوری اطلاعات iostat
func CollectIostatMetrics() {
	cmd := exec.Command("iostat", "-xd", "1", "1")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error running iostat: %v", err)
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	var foundHeader bool
	deviceLineRegex := regexp.MustCompile(`^\S+\s+[\d.]+\s+[\d.]+\s+[\d.]+\s+[\d.]+\s+[\d.]+\s+[\d.]+\s+[\d.]+\s+[\d.]+\s+[\d.]+\s+[\d.]+\s+[\d.]+\s+[\d.]+\s+[\d.]+\s+[\d.]+`)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Device") {
			foundHeader = true
			continue
		}
		if !foundHeader {
			continue
		}

		if deviceLineRegex.MatchString(line) {
			fields := strings.Fields(line)
			if len(fields) < 14 {
				continue
			}

			device := fields[0]
			rps, _ := strconv.ParseFloat(fields[1], 64)
			rkbps, _ := strconv.ParseFloat(fields[2], 64)
			wps, _ := strconv.ParseFloat(fields[7], 64)
			wkbps, _ := strconv.ParseFloat(fields[8], 64)
			util, _ := strconv.ParseFloat(fields[13], 64)

			rrqms, _ := strconv.ParseFloat(fields[3], 64)
			prrqms, _ := strconv.ParseFloat(fields[4], 64)
			rAwait, _ := strconv.ParseFloat(fields[5], 64)
			rareqSz, _ := strconv.ParseFloat(fields[6], 64)
			wrqms, _ := strconv.ParseFloat(fields[9], 64)
			pwrqms, _ := strconv.ParseFloat(fields[10], 64)
			wAwait, _ := strconv.ParseFloat(fields[11], 64)
			wareqSz, _ := strconv.ParseFloat(fields[12], 64)
			avgQueueSize, _ := strconv.ParseFloat(fields[13], 64)

			iostatReadRequests.WithLabelValues(device).Set(rps)
			iostatReadKB.WithLabelValues(device).Set(rkbps)
			iostatWriteRequests.WithLabelValues(device).Set(wps)
			iostatWriteKB.WithLabelValues(device).Set(wkbps)
			iostatUtilization.WithLabelValues(device).Set(util)

			iostatRRQMS.WithLabelValues(device).Set(rrqms)
			iostatPRRQMS.WithLabelValues(device).Set(prrqms)
			iostatRAwait.WithLabelValues(device).Set(rAwait)
			iostatRAreqSz.WithLabelValues(device).Set(rareqSz)
			iostatWRQMS.WithLabelValues(device).Set(wrqms)
			iostatPWRQMS.WithLabelValues(device).Set(pwrqms)
			iostatWAwait.WithLabelValues(device).Set(wAwait)
			iostatWAreqSz.WithLabelValues(device).Set(wareqSz)
			iostatAvgQueueSize.WithLabelValues(device).Set(avgQueueSize)
		}
	}
}
