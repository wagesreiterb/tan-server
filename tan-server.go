package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"math/rand"
	"time"
)

func main() {
	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register functions to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/tan/num", tan_num)
	server.HandleFunc("/tan/anum", tan_anum)

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

func tan_num(w http.ResponseWriter, r *http.Request) {
	log.Printf("func: tan_num; Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 9999
	fmt.Fprintf(w, "Numberic TAN - Version: 1.0.0\n")
	fmt.Fprintf(w, "Serving host: Hostname: %s\n", host)
	fmt.Fprintf(w, "%04d\n", rand.Intn(max - min) + min)
}


// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(65 + rand.Intn(90-65))
    }
    return string(bytes)
}

func tan_anum(w http.ResponseWriter, r *http.Request) {
	log.Printf("func: tan_alpha; serving request: %s", r.URL.Path)
	host, _ := os.Hostname()

	fmt.Fprintf(w, "Alphanumberic TAN - Version: 1.0.0\n")
	fmt.Fprintf(w, "Serving host: Hostname: %s\n", host)
	fmt.Fprintf(w, "%s\n", randomString(6)) // print 10 chars
}
