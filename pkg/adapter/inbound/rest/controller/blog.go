package controller

import (
	"encoding/json"
	"github.com/canmor/go_ms_clean_arch/pkg/app/usecase"
	"github.com/canmor/go_ms_clean_arch/pkg/domain/blog"
	"io"
	"net/http"
)

type Blog struct {
	repo blog.BlogRepository
}

type BlogParam struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type BlogResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewBlog(repo blog.BlogRepository) Blog {
	return Blog{repo}
}

func newResponse(blog *blog.Blog) BlogResponse {
	resp := BlogResponse{int(blog.Id), blog.Title, blog.Body}
	return resp
}

func (b Blog) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var bodyParam BlogParam
	err = json.Unmarshal(body, &bodyParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := usecase.NewBlogUseCase(b.repo)
	created := u.Create(bodyParam.Title, bodyParam.Body)
	if created == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := newResponse(created)
	out, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(out)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
