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

type Test struct {
	Hello string `json:"hello"`
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
	urls := []string{
		"http://localhost:8000",
		"http://localhost:8001",
		"http://localhost:8002",
	}

	minLoad := new(Load)
	for i, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln("Error calling server:", i+1, err)
		}
		defer resp.Body.Close()
		load := new(Load)
		load.ServerId = i + 1
		json.NewDecoder(resp.Body).Decode(&load)
		log.Println(load)

		if i == 0 {
			minLoad = load
		} else {
			minLoad = getMinLoad(minLoad, load)
		}
	}
	log.Println("Min load is:", minLoad)

	resp, err := http.Get(urls[minLoad.ServerId - 1] + "/test")
	if err != nil {
		log.Fatalln("Error calling server:", minLoad.ServerId, err)
	}

	defer resp.Body.Close()

	test := new(Test)
	json.NewDecoder(resp.Body).Decode(&test)
	log.Printf("Test response from min load server %d: %s\n", minLoad.ServerId, test.Hello)
}
