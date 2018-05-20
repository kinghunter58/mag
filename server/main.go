package main

import (
	"fmt"
	"net/http"
	"os"
	// "github.com/gorilla/mux"
)

var PORT = os.Getenv("PORT")

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	fmt.Println("Listening on port:", PORT)

	http.ListenAndServe(":"+PORT, cors(mux))
	// ===== OR with gorilla/mux =====
	// r := mux.NewRouter()
	// r.PathPrefix("/").HandleFunc(indexHandler)
	// http.ListenAndServe(":"+PORT, cors(r))

}

func init() {
	if PORT == "" {
		PORT = "3000"
	}
}
