package main

import (
	"net/http"
	"ss/controller"
)

func main() {
	controller.Start()
	_ = http.ListenAndServe(":8888", nil)
}