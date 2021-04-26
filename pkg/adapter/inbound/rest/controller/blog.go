package controller

import (
	"encoding/json"
	"github.com/canmor/go_ms_clean_arch/pkg/app/usecase"
	"github.com/canmor/go_ms_clean_arch/pkg/domain/blog"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type Blog struct {
	repo     blog.BlogRepository
	shortURL usecase.ShortURL
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

func NewBlog(repo blog.BlogRepository, shortURL usecase.ShortURL) Blog {
	return Blog{repo, shortURL}
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

	u := usecase.NewBlogUseCase(b.repo, b.shortURL)
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

type shareResponse struct {
	Shortcut string `json:"shortcut"`
}

func (b Blog) Share(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	found, err := b.repo.Find(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	use := usecase.NewBlogUseCase(b.repo, b.shortURL)
	shortURL, err := use.Share(*found)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	out, _ := json.Marshal(shareResponse{shortURL})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(out)
}
