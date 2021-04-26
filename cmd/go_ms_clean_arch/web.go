package main

import (
	"database/sql"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/inbound/rest/router"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/outbound/db"
	"net/http"
	"strconv"
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
	return http.ListenAndServe(strconv.Itoa(int(w.config.ListenPort)), w.router)
}
