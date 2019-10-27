package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	port = ":" + os.Getenv("PORT")
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func main() {
	http.HandleFunc("/", hello)
	log.Printf("Listening on %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
