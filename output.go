package main

import (
	"fmt"
)

type KResource struct {
	Memory string `json:"memory"`
	CPU    string `json:"cpu"`
}

type KResources struct {
	Requests KResource `json:"requests"`
}

func ProcessResults(r Resource) KResources {
	result := KResources{
		KResource{
			Memory: convertMemoryToString(r.Memory),
			CPU:    fmt.Sprintf("%0.0fm", r.GetMilliCPU()),
		},
	}
	return result
}

func convertMemoryToString(r int) string {
	if r > 1000 {
		return fmt.Sprintf("%dGi", r)
	}
	return fmt.Sprintf("%dMi", r)
}
