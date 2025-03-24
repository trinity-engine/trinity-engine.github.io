package main

import (
	"path"
	"fmt"
	"net/http"
)

func handler(fs http.FileSystem) http.Handler {
    fileServer := http.FileServer(fs)
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        if _, err := fs.Open(path.Clean(r.URL.Path)); err != nil {
        	http.ServeFile(w, r, "404.html")
        } else {
        	fileServer.ServeHTTP(w, r)
        }
    })
}

func main() {
	http.Handle("/", handler(http.Dir("./")))
	fmt.Println("Starting a server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
