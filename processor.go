package main

import (
	"fmt"
)

func Process(c Config, m []Resource) Resource {
	i := 0

	var pRC Resource
	for _, r := range m {
		if i == 0 {
			pRC = r
		}

		fmt.Println(c)

		if r.CPU > c.CPUPerNode || r.Memory > c.MemoryPerNode {
			break
		}
		pRC = r

		i++
	}

	return pRC
}

func checkDeviation(r1 Resource, r2 Resource, d int) bool {
	df := float32(d)

	cpuPerR1 := r1.CPU / float32(r1.QPS)
	cpuPerR2 := r2.CPU / float32(r2.QPS)
	memPerR1 := float32(r1.Memory / r2.Memory)
	memPerR2 := float32(r2.Memory / r2.Memory)

	fmt.Println(((cpuPerR1 - cpuPerR2) / (cpuPerR1 + cpuPerR2) / 2) * 100)

	if cpuPerR2*df/100 > cpuPerR1 {
		return true
	}

	if memPerR2*df/100 > memPerR1 {
		return true
	}

	return true
}
