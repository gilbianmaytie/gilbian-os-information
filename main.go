package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Get OS information
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Number of CPUs:", runtime.NumCPU())

	// Print CPU usage every second
	for {
		cpuUsage := calculateCPUUsage()
		fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
		time.Sleep(time.Second)
	}
}

func calculateCPUUsage() float64 {
	// Get CPU usage
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	cpuUsage := float64(stat.Sys-stat.HeapReleased) / float64(stat.Sys) * 100
	return cpuUsage
}
