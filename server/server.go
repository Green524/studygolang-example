package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const Name = "a"

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\tname=%v\n", r.URL.Path, strings.Split(r.URL.RawQuery, "=")[1])
}
