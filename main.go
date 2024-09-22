package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <dir> <port>")
		os.Exit(1)
	}

	httpDir := os.Args[1]
	port := os.Args[2]
	fileSystem := http.FileServer(http.Dir(httpDir))
	fmt.Printf("UP Server port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, fileSystem))
}
