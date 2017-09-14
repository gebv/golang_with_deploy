package main

import (
	"fmt"
	"net/http"
	"os"
)

var VERSION = "dev"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version=%s path=%s", VERSION, r.URL.Path)
}

func main() {
	if len(os.Args) < 2 {
		panic("empty address")
	}
	fmt.Println("start listen", os.Args[1])
	http.HandleFunc("/", handler)
	fmt.Println(http.ListenAndServe(os.Args[1], nil))
}
