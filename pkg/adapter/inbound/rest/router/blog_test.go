package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/outbound"
	"github.com/canmor/go_ms_clean_arch/pkg/adapter/outbound/db"
	"github.com/canmor/go_ms_clean_arch/pkg/domain/blog"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func prepareDB(withRecords bool) *sql.DB {
	res, err := db.NewInMemory()
	if err != nil {
		log.Panicf("db error: %s", err)
	}
	err = db.Migrate(res)
	if err != nil {
		log.Panicf("db error: %s", err)
	}
	if !withRecords {
		return res
	}
	repo := outbound.NewBlogRepository(res)
	_, err = repo.Save(blog.Blog{Title: "A blog to share", Body: "body goes here...", CreatedAt: time.Now()})
	if err != nil {
		log.Panicf("db error: %s", err)
	}
	return res
}

func TestBlogCreate(t *testing.T) {
	router := NewRouter(prepareDB(false))
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

func TestBlogShare(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	const shortURL = "https://s.url/j18kNmic2M6"
	responder := httpmock.NewStringResponder(201, fmt.Sprintf(`{"shortcut":%q}`, shortURL))
	httpmock.RegisterResponder(http.MethodPost, "https://api.s.url/short", responder)
	router := NewRouter(prepareDB(true))
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/blogs/1/share", nil)

	router.ServeHTTP(w, req)

	assertions := assert.New(t)
	assertions.Equal(http.StatusOK, w.Result().StatusCode)
	resp := make(map[string]interface{})
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assertions.Equal(shortURL, resp["shortcut"])
}
