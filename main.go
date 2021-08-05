package main

import (
	"fmt"
	"net/http"
	"os"
)

const max = ^uint64(0)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Add("content-type", "text/plain")
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "only POST requests allowed")
		}

		for i := uint64(2); i <= max; i++ {
			fmt.Printf("is %d prime? %t", i, isPrime(i))
		}
	})

	http.ListenAndServe(":"+port, nil)
}

func isPrime(num uint64) bool {
	var divisor uint64
	for divisor = 2; divisor < num; divisor++ {
		if num == divisor {
			continue
		}

		if num%divisor == 0 {
			return false
		}
	}

	return true
}
