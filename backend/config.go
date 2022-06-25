package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pelletier/go-toml/v2"
)

func loadConfig() {

	file, rerr := ioutil.ReadFile("pfs.example.toml")
	if rerr != nil {
		log.Fatal(rerr)
	}

	var cfg Config
	terr := toml.Unmarshal(file, &cfg)
	if terr != nil {
		log.Fatal(terr)
	}

	fmt.Println(cfg.Save_Dir)

	// file, err := os.Open("pfs.example.toml")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file.Close()

	// fmt.Println(file.)
}
