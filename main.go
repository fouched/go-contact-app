package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const port = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "<p>Hello World!</p>")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", Home)

	fmt.Println(fmt.Sprintf("Starting application on %s", port))
	log.Fatalln(http.ListenAndServe(port, nil))
}
