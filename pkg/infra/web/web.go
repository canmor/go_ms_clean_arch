package web

import (
	"database/sql"
	"fmt"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/inbound/rest/router"
	"github.com/canmor/go_ms_clean_arch/pkg/infra/config"
	"github.com/canmor/go_ms_clean_arch/pkg/infra/db"
	"net/http"
)

type Web struct {
	config *config.Config
	db     *sql.DB
	router http.Handler
}

func NewWeb() (*Web, error) {
	cfg := config.New()
	memoryDB, err := db.NewInMemory()
	if err != nil {
		return nil, err
	}
	return &Web{cfg, memoryDB, router.NewRouter(memoryDB)}, nil
}

func (w Web) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", w.config.ListenPort), w.router)
}
