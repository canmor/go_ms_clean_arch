package router

import (
	"database/sql"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/inbound/rest/controller"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/outbound"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewHandler(controller func(w http.ResponseWriter, r *http.Request, _ map[string]string)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		dict := make(map[string]string)
		for _, p := range params {
			dict[p.Key] = p.Value
		}
		controller(w, r, dict)
	}
}

func NewRouter(db *sql.DB) http.Handler {
	router := httprouter.New()
	blog := controller.NewBlog(outbound.NewBlogRepository(db))
	router.POST("/blogs", NewHandler(blog.Create))
	return router
}
