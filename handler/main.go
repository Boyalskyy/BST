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
// @Router /delete [get]

func Remove(t *tree2.Tree) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remove := r.URL.Query().Get("remove")
		xRemove, err := strconv.ParseInt(remove, 10, 64)
		if err != nil {
			fmt.Println(errors.Wrap(err, "convert error"))
			return
		}
		tree2.Delete(t.Root, xRemove)
		fmt.Fprintf(w, "Remove %v from BST", xRemove)
		fmt.Fprintf(w, "\n")
		var s string
		stringBST := tree2.PrintPreOrder(t.Root, &s)
		fmt.Fprintf(w, "Binary Search Tree: %v", stringBST)
		tree2.PrintPreOrder(t.Root, &s)
	})
}

// @Summary Insert
// @Tags insert
// @Description insert value in BST
// @Router /insert [get]

func Insert(t *tree2.Tree) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		insert := r.URL.Query().Get("insert")
		xInsert, err := strconv.ParseInt(insert, 10, 64)
		if err != nil {
			fmt.Println(errors.Wrap(err, "convert error"))
		}
		t.InsertRoot(xInsert)
		fmt.Fprintf(w, "Insert %v to BST", xInsert)
		fmt.Fprintf(w, "\n")
		var s string
		stringBST := tree2.PrintPreOrder(t.Root, &s)
		fmt.Fprintf(w, "Binary Search Tree: %v", stringBST)
		tree2.PrintPreOrder(t.Root, &s)
	})
}

// @Summary Search
// @Tags search
// @Description lookup value in BST
// @Router /search [get]

func Search(t *tree2.Tree) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		xSearch, err := strconv.ParseInt(search, 10, 64)
		if err != nil {
			fmt.Println(errors.Wrap(err, "convert error"))
		}
		searchResult := tree2.Search(t.Root, xSearch)
		fmt.Fprintf(w, "Search %v in BST", xSearch)
		fmt.Fprintf(w, "\n")
		var s string
		stringBST := tree2.PrintPreOrder(t.Root, &s)
		fmt.Fprintf(w, "Binary Search Tree: %v", stringBST)
		tree2.PrintPreOrder(t.Root, &s)
		fmt.Fprintf(w, "Search result %v", searchResult)
	})
}

// @Summary MainPage
// @Tags main
// @Description welcome page
// @Router / [get]

func MainPage(t *tree2.Tree) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Binary Search Tree"))
	})
}
