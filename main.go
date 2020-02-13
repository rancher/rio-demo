package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Arbitrary sleep so that we can demonstrate autoscaler
	time.Sleep(101 * time.Millisecond)
	fmt.Fprintln(w, "Sniper: Test")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
