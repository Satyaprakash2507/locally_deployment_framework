// main.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/Satyaprakash2507/locally_deployment_framework"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!\n")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
