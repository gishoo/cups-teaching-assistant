package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text html; charset=utf-8")
	fmt.Fprint(w, "Home")
}

func voteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text html; charset=utf-8")
	fmt.Fprint(w, "vote")
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text html; charset=utf-8")
	fmt.Fprint(w, "result")
}

func instructorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text html; charset=utf-8")
	fmt.Fprint(w, "instructor")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text html; charset=utf-8")
	fmt.Fprint(w, "admin")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/vote", voteHandler)
	http.HandleFunc("/result", resultHandler)
	http.HandleFunc("/instructor", instructorHandler)
	http.HandleFunc("/admin", adminHandler)

	fmt.Println("Starting Server on Port: 3000")
	http.ListenAndServe(":3000", nil)
}
