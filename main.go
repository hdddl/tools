package main

import (
	"embed"
	"log"
	"net/http"
	"tools/linkProxyDownloader"
)

//go:embed assets/*
var static embed.FS

//go:embed index.html
var index []byte

const addr = "localhost:9000"

func main() {
	http.Handle("/assets/", http.FileServer(http.FS(static)))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(index)
	})
	http.HandleFunc("/linkProxyDownloader/proxy", linkProxyDownloader.ProxyHandler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
