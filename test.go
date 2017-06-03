package main

import (
	"goroute/goroute"
)

func main() {
	router := goroute.New()

	router.Get("/bob/armin")
	router.Get("/bob/bill")
	router.Get("/armin/xun")
	router.Get("/armin/bob/:id")
	router.Get("/armin/bob/:id/zoo")
	router.Get("/armin/bob/noob")
	router.Get("/armin/bob/noob/:name")

	//router.PrintTrees()

	router.Call("GET", "/bob/armin")
	router.Call("GET", "/armin/bob/5/zoo")
	router.Call("GET", "/armin/bob/noob/xun")

}
