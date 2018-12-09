package main

func Process(c Config, m []Resource) Resource {
	i := 0

	var pRC Resource
	for _, r := range m {
		if i == 0 {
			pRC = r
		}

		if r.CPU > c.CPUPerNode || r.Memory > c.MemoryPerNode || checkDeviation(pRC, r, c.Deviation) {
			break
		}
		pRC = r

		i++
	}

	return pRC
}

func checkDeviation(r1 Resource, r2 Resource, d float64) bool {
	cpuPerR1 := r1.CPU / float64(r1.QPS)
	cpuPerR2 := r2.CPU / float64(r2.QPS)
	memPerR1 := float64(r1.Memory / r1.QPS)
	memPerR2 := float64(r2.Memory / r2.QPS)

	comp := func(a float64, b float64, d float64) bool {
		return 100-100*a/b > d
	}

	if comp(cpuPerR1, cpuPerR2, d) {
		return true
	}

	if comp(memPerR1, memPerR2, d) {
		return true
	}

	return false
}
