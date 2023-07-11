package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
)

type Load struct {
	elapsed float64
	cpu float64
	memory float64
	disk float64
}

func measureLoad(url string) (*Load) {
	start := time.Now()

	// Measure response time
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	elapsed := time.Since(start).Seconds()

	// Measure CPU load
	percent, _ := cpu.Percent(time.Second, false)

	// Measure memory usage
	v, _ := mem.VirtualMemory()

	// Measure disk usage
	d, _ := disk.Usage("/")

	// return elapsed, percent[0], v.UsedPercent, d.UsedPercent
	return &Load{
		elapsed: elapsed,
		cpu: percent[0],
		memory: v.UsedPercent,
		disk: d.UsedPercent,
	}
}

func main() {
	server1Load := measureLoad("http://localhost:8000")
	server2Load := measureLoad("http://localhost:8000")
	server3Load := measureLoad("http://localhost:8000")

	fmt.Println(server1Load, server2Load, server3Load)
}
