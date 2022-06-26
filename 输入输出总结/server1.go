package main

import (
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receive a request from:", r.RemoteAddr, r.Header)
	w.Write([]byte("ok"))
}

// net/http 默认开启持久连接
func main() {
	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", http.HandlerFunc(Index)); err != nil {
		log.Fatal(err)
	}
}
