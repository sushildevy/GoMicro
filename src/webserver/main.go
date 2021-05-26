package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("Hello go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
