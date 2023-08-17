package main

import "net/http"

func main() {
	roothandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./../front-end/index.html")
	}

	http.HandleFunc("/", roothandler)

	http.ListenAndServe("127.0.0.1:9000", nil)
}
