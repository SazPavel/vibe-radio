// vibe-only — автономный сервер одной страницы vibe.html.
//
// Никакого функционала тестов: только раздача встроенной страницы.
// Запуск: go run . [-addr :8090]
package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static/vibe.html
var static embed.FS

func main() {
	addr := flag.String("addr", ":8090", "адрес слушающего сокета")
	flag.Parse()

	sub, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatal(err)
	}
	page, err := fs.ReadFile(sub, "vibe.html")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	serve := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(page)
	}
	mux.HandleFunc("/", serve) // любой путь — одна и та же страница

	log.Printf("vibe-only: http://localhost%s", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}
