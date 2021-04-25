package main

import (
	"log"
)

func main() {
	w, err := NewWeb()
	if err != nil {
		log.Panicf("error: %s", err)
	}
	err = w.Run()
	if err != nil {
		log.Panicf("error: %s", err)
	}
}
