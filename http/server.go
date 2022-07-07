package http

import (
	"fmt"
	"net/http"
)

func Server() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}
