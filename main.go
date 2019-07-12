package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func findRecord(q string) (url string, err error) {
	lines, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer lines.Close()
	r := csv.NewReader(lines)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == q {
			return record[1], nil
		}
	}
	return "", nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	uri, err := findRecord(code)
	if err != nil {
		log.Fatal(err)
	}
	if uri == "" {
		http.NotFound(w, r)
	}
	fmt.Println(code)
	http.Redirect(w, r, uri, 301)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
