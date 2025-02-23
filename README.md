# iostat_exporter
🚀 **A Prometheus exporter for monitoring disk I/O statistics using iostat.**  
This exporter collects disk I/O metrics and exposes them in **Prometheus-compatible format**.
Exposes metrics such as read/write requests, disk utilization, and queue size.

## 📝 About the Project  
The `iostat_exporter` is a lightweight Prometheus exporter for monitoring disk I/O performance using `iostat`.  
This exporter is useful in environments where disk performance monitoring is critical, such as:
- **Database servers** (e.g., PostgreSQL, MySQL, MongoDB)
- **High IOPS applications**
- **Cloud and virtualized environments**

---

## 📦 Installation

### 🔹 Prerequisites
- `iostat` (from `sysstat`)
- `Go 1.20+`
- `Prometheus` (optional)
- `systemd` (for running as a service)

### 🏗️ Build from Source  
#### **Linux/macOS**
```sh
git clone https://github.com/nikousokhan/iostat_exporter.git
cd iostat_exporter
go build -o iostat_exporter main.go
```

## Run the Exporter
```sh
./iostat_exporter
```
- By default, runs on :9100 → http://localhost:9100/metrics.

By default, it will listen on :9100. You can change the host and port via environment variables:
```sh
EXPORTER_HOST=0.0.0.0 EXPORTER_PORT=9200 ./iostat_exporter
```

## 📊 Exported Metrics

## Exported Metrics

| **Metric**               | **iostat Equivalent** | **Description**                     |
|--------------------------|----------------------|-------------------------------------|
| `iostat_read_requests`   | `r/s`               | Read requests per second           |
| `iostat_read_kb_per_sec` | `rkB/s`             | KB read per second                 |
| `iostat_write_requests`  | `w/s`               | Write requests per second          |
| `iostat_write_kb_per_sec`| `wkB/s`             | KB written per second              |
| `iostat_avg_queue_size`  | `aqu-sz`            | Avg. I/O requests in queue         |
| `iostat_utilization`     | `%util`             | Disk utilization (%)               |
| `iostat_r_await`        | `r_await`           | Read request wait time (ms)        |
| `iostat_w_await`        | `w_await`           | Write request wait time (ms)       |
| `iostat_rareq_sz`       | `rareq-sz`          | Avg. read request size (KB)        |
| `iostat_wareq_sz`       | `wareq-sz`          | Avg. write request size (KB)       |

## ⚙️ Running as a Systemd Service

To run `iostat_exporter` as a **systemd service**, follow these steps:

### 1️⃣ Create the service file

- Create a new systemd service file:

```sh
sudo nano /etc/systemd/system/iostat_exporter.service
```
- Add the following content:
```sh
[Unit]
Description=Iostat Exporter for Prometheus
After=network.target

[Service]
User=root
ExecStart=/usr/local/bin/iostat_exporter
Restart=always
Environment="EXPORTER_HOST=0.0.0.0"
Environment="EXPORTER_PORT=9200"

[Install]
WantedBy=multi-user.target
```
### 2️⃣ Reload systemd and enable the service
- Run the following commands:
```sh
sudo systemctl daemon-reload
sudo systemctl enable iostat_exporter
sudo systemctl start iostat_exporter
```

### 3️⃣ Check service status
- Verify that the service is running:
```sh
sudo systemctl status iostat_exporter
```
## 🔧 Configuration

You can configure the **host** and **port** of `iostat_exporter` via the systemd service file.

### 1️⃣ Create the systemd service file

Run:

```sh
sudo nano /etc/systemd/system/iostat_exporter.service
```
You can change this part of unit
- Add the following content:
```sh
Environment="EXPORTER_HOST=0.0.0.0"
Environment="EXPORTER_PORT=9200"
```

### 2️⃣ Reload systemd and restart the service
```sh
sudo systemctl daemon-reload
sudo systemctl restart iostat_exporter
```
Now, the exporter will run on 0.0.0.0:9200 instead of the default port.

To verify:
```sh
sudo systemctl status iostat_exporter
```
- For Prometheus integration, add this to prometheus.yml:
```sh 
scrape_configs:
  - job_name: 'iostat_exporter'
    static_configs:
      - targets: ['localhost:9100']
```
## 🛠️ Development & Contribution
Clone the repo and start contributing!
```sh 
git clone https://github.com/nikousokhan/iostat_exporter.git
```
## 🩺 How to Add a Health Check for the Exporter?

Since your exporter runs an HTTP server and exposes metrics at `/metrics`, you should add a **simple health check** using `curl` to verify that it's working properly.

### Test the Exporter Manually  

After starting the exporter, you can check if the metrics are available by running:  

```sh
curl -s http://localhost:9100/metrics | head -n 20
```