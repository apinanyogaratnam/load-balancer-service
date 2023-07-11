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
}


func main() {
	resp, err := http.Get("http://localhost:8001")
	if err != nil {
		log.Fatalln("Error calling server 1")
	}
	defer resp.Body.Close()
	load := new(Load)
	json.NewDecoder(resp.Body).Decode(&load)
	log.Println(load)

	resp1, err := http.Get("http://localhost:8002")
	if err != nil {
		log.Fatalln("Error calling server 2")
	}
	defer resp1.Body.Close()
	load1 := new(Load)
	json.NewDecoder(resp1.Body).Decode(&load1)
	log.Println(load1)

	resp2, err := http.Get("http://localhost:8003")
	if err != nil {
		log.Fatalln("Error calling server 3")
	}
	defer resp2.Body.Close()
	load2 := new(Load)
	json.NewDecoder(resp2.Body).Decode(&load2)
	log.Println(load2)
}
