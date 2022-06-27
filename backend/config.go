package main

import (
	"fmt"
	"io/ioutil"

	"github.com/pelletier/go-toml/v2"
)

func ConvertConfigs(pc *PreConfig) *Config {
	return &Config{
		Save_Dir:    pc.Save_Dir,
		Expiry:      int64(pc.Expiry),
		Max_Storage: int64(pc.Max_Storage),
	}
}

func LoadConfig() (*Config, error) {

	file, rerr := ioutil.ReadFile("pfs.example.toml")
	if rerr != nil {
		return nil, rerr
	}

	var pcfg PreConfig
	terr := toml.Unmarshal(file, &pcfg)
	if terr != nil {
		return nil, terr
	}

	cfg := ConvertConfigs(&pcfg)

	fmt.Println(cfg.Save_Dir)
	return cfg, nil
}
