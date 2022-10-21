package router

import (
	"BST/handler"
	tree2 "BST/tree"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pseidemann/finish"
	_ "github.com/pseidemann/finish"
	"log"
	"net/http"
	"time"
)

func NewRouter(t tree2.Tree) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handler.MainPage(&t))
	r.Get("/delete", handler.Remove(&t))
	r.Get("/insert", handler.Insert(&t))
	r.Get("/search", handler.Search(&t))
	srv := &http.Server{Addr: ":8080", Handler: r}
	fin := &finish.Finisher{Timeout: 10 * time.Second}
	fin.Add(srv)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	fin.Wait()
}
