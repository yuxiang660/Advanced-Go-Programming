package main

import (
	"io/ioutil"
	"net/http"
	"log"
)

func echo(res http.ResponseWriter, req *http.Request) {
	msg, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	writeLen, err := res.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
