package main

import "net/http"

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("<h1>Hello server</h1>"))
	})

	http.ListenAndServe(":8080", nil)
}
