package main

import (
	"os"
	"strconv"
)

type Config struct {
	ListenPort int16
}

func New() *Config {
	result := Config{ListenPort: 8080}
	listenPort, ok := os.LookupEnv("LISTEN_PORT")
	if ok {
		port, err := strconv.Atoi(listenPort)
		if err != nil {
			result.ListenPort = int16(port)
		}
	}
	return &result
}
