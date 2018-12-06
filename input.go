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

func readDockerStat(file string) dockerStat {
	var ds dockerStat

	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Can't read file :", file, "Error : ", err)
	} else {
		json.Unmarshal(f, &ds)
		fmt.Println(ds)
	}
	return ds
}

func readVegeta(file string) dockerStat {
	var ds dockerStat

	f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Can't read file :", file, "Error : ", err)
	} else {
		json.Unmarshal(f, &ds)
		fmt.Println(ds)
	}
	return ds
}
