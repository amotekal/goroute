package main

import (
	"fmt"
	"goroute/goroute"
	"net/http"
)

func testhandle(w http.ResponseWriter, r *http.Request) {
	id := goroute.Param(r, "id")
	fmt.Fprintln(w, "hello world"+id)
}

func panichandler(w http.ResponseWriter, r *http.Request, rcv interface{}) {
	fmt.Fprintln(w, rcv)
}

func main() {
	router := goroute.New()
	router.Get("/armin/bob/:id", testhandle)
	router.PanicHandler = panichandler
	//router.PrintTrees()
	http.ListenAndServe(":8080", router)

}
