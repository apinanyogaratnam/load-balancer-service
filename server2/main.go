package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type Load struct {
	Cpu     float64 `json:"cpu"`
	Memory  float64 `json:"memory"`
	Disk    float64 `json:"disk"`
}

func measureLoad() *Load {
	// Measure response time
	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	// Measure CPU load
	percent, _ := cpu.Percent(0, false)

	// Measure memory usage
	v, _ := mem.VirtualMemory()

	// Measure disk usage
	d, _ := disk.Usage("/")

	return &Load{
		Cpu:     percent[0],
		Memory:  v.UsedPercent,
		Disk:    d.UsedPercent,
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(measureLoad())
		log.Println("Request received")
	})

	http.ListenAndServe(":8002", r)
}
