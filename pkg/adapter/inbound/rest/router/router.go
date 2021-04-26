package router

import (
	"database/sql"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/inbound/rest/controller"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/outbound"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(db *sql.DB) http.Handler {
	router := mux.NewRouter()
	blog := controller.NewBlog(outbound.NewBlogRepository(db))
	router.HandleFunc("/blogs", blog.Create)
	return router
}
