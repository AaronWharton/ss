package main

import (
	"log"
	"net/http"
	"ss/controller"
)

func main() {
	controller.Startup()
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln(err)
	}
}