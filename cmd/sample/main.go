package main

import (
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	Port string
	Name string
}

var config Config = Config{
	Port: "8080",
	Name: "A",
}

func init() {
	fmt.Println("Server is starting...")
	name := os.Getenv("NAME")
	if name != "" {
		config.Name = name
	}
	port := os.Getenv("PORT")
	if port != "" {
		config.Port = port
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Microservice! %s", config.Name)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)
	fmt.Printf("Server %s berjalan di port %s", config.Name, config.Port)
	serverPort := fmt.Sprintf(":%s", config.Port)
	http.ListenAndServe(serverPort, nil)
}
