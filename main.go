package main

import (
	"BST/router"
	tree2 "BST/tree"
	_ "fmt"
	_ "github.com/pkg/errors"
	_ "strconv"
)

// @title BST
// @version 1.0
// @description API Server for Search,Insert,Delete value in BST

// @host localhost:8080
// @BasePath /

func main() {
	tree := tree2.NewTree()
	_ = router.NewRouter(tree)

}
