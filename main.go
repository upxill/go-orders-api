package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello API")
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to start server: %v", err)
	}

}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
