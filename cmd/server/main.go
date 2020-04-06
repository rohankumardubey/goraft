package main

import (
	"fmt"
	"log"
	"os"
	"path"

	raftconfig "github.com/amukherj/raft/internal/config"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panicf("Could not determine working directory: %v", err)
	}

	yamlConfig := path.Join(cwd, "raft.yaml")
	configReader := raftconfig.NewFileConfigReader(yamlConfig)
	config, err := configReader.Read()
	if err != nil {
		log.Printf("Could not read raft config from %s: %v", yamlConfig, err)
		jsonConfig := path.Join(cwd, "raft.json")
		configReader = raftconfig.NewFileConfigReader(yamlConfig)
		config, err = configReader.Read()
		if err != nil {
			log.Panicf("Could not read raft config from %s: %v", jsonConfig, err)
		}
	}

	fmt.Printf("Config read: %+v", config)
}
