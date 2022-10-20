package router

import (
	_ "BST/docs"
	"BST/handler"
	tree2 "BST/tree"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func NewRouter(t tree2.Tree) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handler.MainPage(&t))
	r.Get("/delete", handler.Remove(&t))
	r.Get("/insert", handler.Insert(&t))
	r.Get("/search", handler.Search(&t))
	r.Mount("/swagger", httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", r)
	return r
}
