package main

// # --- Configuration file ---
import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	NodeNumber    int     `yaml:"nodeNumber"`
	CPUPerNode    float32 `yaml:"cpuPerNode"`
	MemoryPerNode int     `yaml:"memoryPerNode"`

	Deviation float32 `yaml:"devivation"`
}

func ParseConfig(file string) Config {
	var c Config

	f, err := ioutil.ReadFile(file)
	check(err)

	err = yaml.Unmarshal(f, &c)
	check(err)

	return c
}
