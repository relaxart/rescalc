package main

import (
	"encoding/json"
	"os"

	tm "github.com/buger/goterm"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	tm.Clear()
	tm.Println("ResCal processing ...")

	c := ParseConfig("./config.yaml")
	tm.Printf("Resource deviation: %.0f%%\n", c.Deviation)
	result := os.Args[1:]
	if len(result) < 1 {
		tm.Println("You must pass results")
	} else {
		r := readResult(result[0])
		p := Process(c, r.Results)
		o := ProcessResults(p)

		json, err := json.MarshalIndent(o, "", "\t")
		check(err)

		cn := r.QPS / p.QPS
		tm.Printf("Recommended number of containers: %d\n", cn)
		tm.Println("Calculated resources:")
		tm.Println(string(json))

		chartM := tm.NewLineChart(80, 15)
		chartM.Flags = tm.DRAW_RELATIVE

		dataM := new(tm.DataTable)
		dataM.AddColumn("QPS")
		dataM.AddColumn("Memory")

		for _, r := range r.Results {
			dataM.AddRow(float64(r.QPS), float64(r.Memory))
		}

		tm.Println(chartM.Draw(dataM))
		tm.Flush()

		chartC := tm.NewLineChart(80, 15)
		chartC.Flags = tm.DRAW_RELATIVE
		dataC := new(tm.DataTable)
		dataC.AddColumn("QPS")
		dataC.AddColumn("CPU")

		for _, r := range r.Results {
			dataC.AddRow(float64(r.QPS), r.CPU)
		}

		tm.Println(chartC.Draw(dataC))
		tm.Flush()
	}
}
