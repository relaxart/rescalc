package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type memory struct {
	Raw     string
	Rercent string
}

type dockerStat struct {
	Container string
	CPU       string
	Memory    memory
}

type latency struct {
	Total int
	Mean  int
	P50th int `json:"50th"`
	P95th int `json:"95th"`
	P99th int `json:"99th"`
	Max   int
}

type vegeta struct {
	Requests  int
	Latencies latency
}

func readDockerStat(file string) dockerStat {
	var ds dockerStat

	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Can't read file :", file, "Error : ", err)
	} else {
		json.Unmarshal(f, &ds)
	}
	return ds
}

func readVegeta(file string) vegeta {
	var v vegeta

	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Can't read file :", file, "Error : ", err)
	} else {
		json.Unmarshal(f, &v)
	}
	return v
}
