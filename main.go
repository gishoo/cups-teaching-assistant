package main

import (
	"fmt"
	"net/http"
)

func routeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "home")
	case "/vote":
		fmt.Fprint(w, "vote")
	case "/result":
		fmt.Fprint(w, "result")
	case "/instructor":
		fmt.Fprint(w, "instructor")
	case "/admin":
		fmt.Fprint(w, "admin")
	}
}

func main() {
	http.HandleFunc("/", routeHandler)

	fmt.Println("Starting Server on Port: 3000")
	http.ListenAndServe(":3000", nil)
}
