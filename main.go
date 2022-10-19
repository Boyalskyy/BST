package main

import (
	"BST/router"
	tree2 "BST/tree"
	_ "fmt"
	_ "github.com/pkg/errors"
	_ "strconv"
)

func main() {
	tree := tree2.NewTree()
	_ = router.NewRouter(tree)

}

//r := chi.NewRouter()
//r.Use(middleware.Logger)
//r.Get("/", handler.MainPage(&tree))
//r.Get("/delete", handler.Remove(&tree))
//r.Get("/insert", handler.Insert(&tree))
//r.Get("/search", handler.Search(&tree))
//http.ListenAndServe(":8080", r)