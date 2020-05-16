package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Main function
func main() {

	r := mux.NewRouter()
	// Route handles & endpoints
	r.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		// json, _ := json.Marshal("Welcome home")

		var res RESP
		err := json.NewDecoder(r.Body).Decode(&res)
		if err != nil {
			w.Write([]byte("error"))
			log.Println(err)
			// return
		}

		// log.Println(res)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		buf := &bytes.Buffer{}
		json.NewEncoder(buf).Encode(&res)

		w.Write(buf.Bytes())

	}).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":8000", r)
}

type RESP struct {
	Value string `json:"value"`
}
