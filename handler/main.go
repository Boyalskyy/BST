package handler

import (
	tree2 "BST/tree"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

// @Summary Remove
// @Tags remove
// @Description remove value from BST
// @Param remove query int true "int"
// @Router /delete [get]
func Remove(t *tree2.Tree) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		remove := r.URL.Query().Get("remove")
		xRemove, err := strconv.ParseInt(remove, 10, 64)
		if err != nil {
			fmt.Println(errors.Wrap(err, "convert error"))
			return
		}
		tree2.Delete(t.Root, xRemove)
		fmt.Fprintf(w, "\n")
		var s string
		stringBST := tree2.PrintPreOrder(t.Root, &s)
		fmt.Fprintf(w, "Remove %v from BST", xRemove)
		fmt.Fprintf(w, "Binary Search Tree: %v", stringBST)
	}
}

// @Summary Insert
// @Tags insert
// @Description insert value in BST
// @Param insert query int true "int"
// @Router /insert [get]
func Insert(t *tree2.Tree) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		insert := r.URL.Query().Get("insert")
		xInsert, err := strconv.ParseInt(insert, 10, 64)
		if err != nil {
			fmt.Println(errors.Wrap(err, "convert error"))
		}
		t.InsertRoot(xInsert)
		var s string
		stringBST := tree2.PrintPreOrder(t.Root, &s)
		fmt.Fprintf(w, "Insert %v to BST", xInsert)
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "Binary Search Tree: %v", stringBST)
	}
}

// @Summary Search
// @Tags search
// @Description lookup value in BST
// @Param search query int true "int"
// @Router /search [get]
func Search(t *tree2.Tree) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		xSearch, err := strconv.ParseInt(search, 10, 64)
		if err != nil {
			fmt.Println(errors.Wrap(err, "convert error"))
		}
		searchResult := tree2.Search(t.Root, xSearch)
		var s string
		stringBST := tree2.PrintPreOrder(t.Root, &s)
		fmt.Fprintf(w, "Search %v in BST", xSearch)
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "Binary Search Tree: %v", stringBST)
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "Search result %v", searchResult)
	}
}

// @Summary MainPage
// @Tags main
// @Description welcome page
// @Router / [get]
func MainPage(t *tree2.Tree) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Binary Search Tree"))
	}
}
