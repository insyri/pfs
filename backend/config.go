package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

func ConvertConfigs(pc *PreConfig) *Config {
	return &Config{
		Save_Dir:    pc.Save_Dir,
		Expiry:      int64(pc.Expiry),
		Max_Storage: int64(pc.Max_Storage),
		Db_User:     pc.Db_User,
		Db_Pass:     pc.Db_Pass,
		Db_Name:     pc.Db_Name,
	}
}

func LoadConfig() (*Config, error) {

	// Read TOML file
	file, rerr := ioutil.ReadFile("pfs.example.toml")
	if rerr != nil {
		return nil, rerr
	}

	// Extract TOML entries to PreConfig struct
	var pcfg PreConfig
	terr := toml.Unmarshal(file, &pcfg)
	if terr != nil {
		return nil, terr
	}

	// Convert PreConfig struct to Config struct (thanks, TOML)
	cfg := ConvertConfigs(&pcfg)

	return cfg, nil
}
