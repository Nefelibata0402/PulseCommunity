package new_prometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func InitPrometheus() {
	go func() {
		fmt.Println("Starting Prometheus metrics server on :8083") // 添加日志
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":8081", nil); err != nil {
			fmt.Println("Error starting metrics server:", err)
		}
	}()
}

func InitPrometheusUser() {
	go func() {
		fmt.Println("Starting Prometheus metrics server on :8083") // 添加日志
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":8082", nil); err != nil {
			fmt.Println("Error starting metrics server:", err)
		}
	}()
}

func InitPrometheusArticle() {
	go func() {
		fmt.Println("Starting Prometheus metrics server on :8083") // 添加日志
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":8083", nil); err != nil {
			fmt.Println("Error starting metrics server:", err)
		}
	}()
}

func InitPrometheusRanking() {
	go func() {
		fmt.Println("Starting Prometheus metrics server on :8084") // 添加日志
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":8084", nil); err != nil {
			fmt.Println("Error starting metrics server:", err)
		}
	}()
}
