package main

import (
	"github.com/BurntSushi/toml"
	"time"

	"fmt"
)

func main() {

	var tomlData = `Age = 25
Cats = [ "Cauchy", "Plato" ]
Pi = 3.14
Perfection = [ 6, 28, 496, 8128 ]
DOB = 1987-07-05T05:45:00Z`
	type Config struct {
		Age        int
		Cats       []string
		Pi         float64
		Perfection []int
		DOB        time.Time
	}
	var conf Config
	_, err := toml.Decode(tomlData, &conf)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(conf.Age)
	//fmt.Print(err)

}
