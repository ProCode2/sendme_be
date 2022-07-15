package main

import (
	"log"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	port := ":" + getEnv("PORT", "8080")
	wsServer := NewWebsocketServer()

	// run websocket server in a go routine
	go wsServer.Run()

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
