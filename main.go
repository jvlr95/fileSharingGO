package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "$ADMIN" {
		return "$PASS"
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <dir> <port>")
		os.Exit(1)
	}

	httpDir := os.Args[1]
	port := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("myserver.com", Secret)

	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))

	fmt.Printf("UP Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
