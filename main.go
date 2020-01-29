package main

import (
	"log"
	"net/http"
	"webserv/handler"
)

func main() {
	http.HandleFunc("/", handler.SayHello)
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
