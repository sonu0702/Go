package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/header", headers)
	fmt.Println("This is connecting")
	err := http.ListenAndServe(":8081", nil)
	fmt.Println(err)
}
