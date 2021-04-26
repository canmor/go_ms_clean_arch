package router

import (
	"database/sql"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/inbound/rest/controller"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/outbound"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/outbound/gateway"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(db *sql.DB) http.Handler {
	router := mux.NewRouter()
	blog := controller.NewBlog(outbound.NewBlogRepository(db), gateway.NewShortURLGatewayImpl())
	router.HandleFunc("/blogs", blog.Create).Methods(http.MethodPost)
	router.HandleFunc("/blogs/{id}/share", blog.Share).Methods(http.MethodGet)
	return router
}
