package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

type OSInfo struct {
	OS         string `json:"os"`
	Arch       string `json:"arch"`
	NumCPU     int    `json:"numCPU"`
	CPUUsage   string `json:"cpuUsage"`
	GoVersion  string `json:"goVersion"`
	NumThreads int    `json:"numThreads"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		info := OSInfo{
			OS:         runtime.GOOS,
			Arch:       runtime.GOARCH,
			NumCPU:     runtime.NumCPU(),
			CPUUsage:   calculateCPUUsage(),
			GoVersion:  runtime.Version(),
			NumThreads: runtime.NumGoroutine(),
		}
		jsonResponse, err := json.Marshal(info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func calculateCPUUsage() string {
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	cpuUsage := float64(stat.Sys-stat.HeapReleased) / float64(stat.Sys) * 100
	return fmt.Sprintf("%.2f%%", cpuUsage)
}
