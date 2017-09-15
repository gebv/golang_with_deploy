package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var VERSION = "dev"

func handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"verions": VERSION,
		"path":    r.URL.Path,
		"status":  "ok",
	})
}

func main() {
	if len(os.Args) < 2 {
		panic("empty address")
	}
	fmt.Println("start listen", os.Args[1])
	http.HandleFunc("/", handler)
	fmt.Println(http.ListenAndServe(os.Args[1], nil))
}
