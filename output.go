package main

import (
	"encoding/json"
	"fmt"
)

type KResource struct {
	Memory string `json:"memory"`
	CPU    string `json:"cpu"`
}

type KResources struct {
	Requests KResource `json:"requests"`
}

func ProcessResults(r Resource) {
	result := KResources{
		KResource{
			Memory: convertMemoryToString(r.Memory),
			CPU:    fmt.Sprintf("%0.0fm", r.GetMilliCPU()),
		},
	}

	json, err := json.MarshalIndent(result, "", "\t")
	check(err)

	fmt.Println(string(json))
}

func convertMemoryToString(r int) string {
	if r > 1000 {
		return fmt.Sprintf("%dGi", r)
	}
	return fmt.Sprintf("%dMi", r)
}
