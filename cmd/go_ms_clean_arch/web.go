package main

import (
	"database/sql"
	"fmt"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/inbound/rest/router"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/outbound/db"
	"net/http"
)

type Web struct {
	config *Config
	db     *sql.DB
	router http.Handler
}

func NewWeb() (*Web, error) {
	cfg := New()
	memoryDB, err := db.NewInMemory()
	if err != nil {
		return nil, err
	}
	return &Web{cfg, memoryDB, router.NewRouter(memoryDB)}, nil
}

func (w *Web) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", w.config.ListenPort), w.router)
}
