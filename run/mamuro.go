package main

import (
	"basededatos/backend"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("\x1b[31;1mError: Missing arguments\x1b[0m\n")
		return
	}

	if os.Args[1] != "-p" || os.Args[2] == "" {
		fmt.Printf("\x1b[31;1mError: Missing port number e.g ./mamuro -p 8080\x1b[0m\n")
		return
	}

	port := os.Args[2]

	go backend.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("../mamuro-email/dist")).ServeHTTP(w, r)
	})

	fmt.Printf("\x1b[32;1mServer listening on port %s...\x1b[0m\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
