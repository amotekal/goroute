package main

import "goroute/goroute"

func main() {
	router := goroute.New()

	router.Get("armin/xun/bill")
	router.Get("armin/xun/bob")

	router.PrintTrees()
}
