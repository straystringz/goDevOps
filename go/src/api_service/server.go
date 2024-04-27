package main

import "net/http"

func main() {
	http.HandleFunc("/hello-world", handleHelloWorld)
}
