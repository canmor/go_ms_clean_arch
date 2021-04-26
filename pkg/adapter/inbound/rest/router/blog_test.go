package router

import (
	"database/sql"
	"encoding/json"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/outbound/db"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func prepareDB() *sql.DB {
	res, err := db.NewInMemory()
	if err != nil {
		log.Panicf("db error: %s", err)
	}
	err = db.Migrate(res)
	if err != nil {
		log.Panicf("db error: %s", err)
	}
	return res
}

func TestBlogCreate(t *testing.T) {
	router := NewRouter(prepareDB())
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/blogs", strings.NewReader(`{"title":"test", "body":"body"}`))

	router.ServeHTTP(w, req)

	assertions := assert.New(t)
	assertions.Equal(http.StatusCreated, w.Result().StatusCode)
	resp := make(map[string]interface{})
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assertions.Equal(1.0, resp["id"])
	assertions.Equal("test", resp["title"])
	assertions.Equal("body", resp["body"])
}
