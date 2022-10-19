package router

import (
	"BST/handler"
	tree2 "BST/tree"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewRouter(t tree2.Tree) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handler.MainPage(&t))
	r.Get("/delete", handler.Remove(&t))
	r.Get("/insert", handler.Insert(&t))
	r.Get("/search", handler.Search(&t))
	http.ListenAndServe(":8080", r)
	return r
}
