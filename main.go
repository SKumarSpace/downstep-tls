package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "9000", "port to listen on")
	flag.Parse()

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		url := queryParams.Get("url")
		if url == "" {
			http.Error(w, "Missing url query parameter", http.StatusBadRequest)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s", body)
	})

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
