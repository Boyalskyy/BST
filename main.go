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
