package main

import (
	"encoding/json"
	"net/http"
	"log"
)

type Load struct {
	Cpu     float64 `json:"cpu"`
	Memory  float64 `json:"memory"`
	Disk    float64 `json:"disk"`
	ServerId int    `json:"server_id"`
}

func (load Load) sum() float64 {
	return load.Cpu + load.Memory + load.Disk
}

func getMinLoad(load1 *Load, load2 *Load) *Load {
	if load1.sum() < load2.sum() {
		return load1
	}

	return load2
}

func main() {
	resp, err := http.Get("http://localhost:8001")
	if err != nil {
		log.Fatalln("Error calling server 1")
	}
	defer resp.Body.Close()
	load := new(Load)
	load.ServerId = 1
	json.NewDecoder(resp.Body).Decode(&load)
	log.Println(load)
	minLoad := load

	resp1, err := http.Get("http://localhost:8002")
	if err != nil {
		log.Fatalln("Error calling server 2")
	}
	defer resp1.Body.Close()
	load1 := new(Load)
	load1.ServerId = 2
	json.NewDecoder(resp1.Body).Decode(&load1)
	log.Println(load1)
	minLoad = getMinLoad(minLoad, load1)

	resp2, err := http.Get("http://localhost:8003")
	if err != nil {
		log.Fatalln("Error calling server 3")
	}
	defer resp2.Body.Close()
	load2 := new(Load)
	load2.ServerId = 3
	json.NewDecoder(resp2.Body).Decode(&load2)
	log.Println(load2)
	minLoad = getMinLoad(minLoad, load2)

	log.Println("min load", minLoad)
}
