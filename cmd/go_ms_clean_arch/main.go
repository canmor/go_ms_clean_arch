package main

import (
	"github.com/canmor/go_ms_clean_arch/pkg/infra/web"
	"log"
)

func main() {
	w, err := web.NewWeb()
	if err != nil {
		log.Panicf("error: %s", err)
	}
	err = w.Run()
	if err != nil {
		log.Panicf("error: %s", err)
	}
}
