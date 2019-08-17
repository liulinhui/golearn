package fileServer

import "net/http"

func FileServer() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	_ = http.ListenAndServe(":8080", nil)
}
