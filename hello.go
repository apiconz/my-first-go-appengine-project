package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
