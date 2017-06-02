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

	router.PrintTrees()

}
