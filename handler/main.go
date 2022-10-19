package handler

import (
	tree2 "BST/tree"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

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
		tree2.PrintPreOrder(t.Root)
	})
}
func Insert(t *tree2.Tree) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		insert := r.URL.Query().Get("insert")
		xInsert, err := strconv.ParseInt(insert, 10, 64)
		if err != nil {
			fmt.Println(errors.Wrap(err, "convert error"))
		}
		t.InsertRoot(xInsert)
		fmt.Fprintf(w, "Insert %v to BST", xInsert)
		tree2.PrintPreOrder(t.Root)
	})
}
func Search(t *tree2.Tree) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")
		xSearch, err := strconv.ParseInt(search, 10, 64)
		if err != nil {
			fmt.Println(errors.Wrap(err, "convert error"))
		}
		tree2.Search(t.Root, xSearch)
		fmt.Fprintf(w, "Search %v in BST", xSearch)
		tree2.PrintPreOrder(t.Root)
	})
}
func MainPage(t *tree2.Tree) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Binary Search Tree"))
	})
}
