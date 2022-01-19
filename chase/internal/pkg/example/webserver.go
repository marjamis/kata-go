package example

import (
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Custom Header", "Here I am")
	log.Printf("Incoming request...\n")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World\n"))
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request...\n")
	w.Write([]byte("Goodbye\n"))
}

func runWebserver() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/goodbye", goodbye)
	http.ListenAndServe(":8000", nil)
}

func init() {
	examples := runs{
		{"Basic Webserver", runWebserver},
	}
	GetMyExamples().Add("webserver", examples.runExamples)
}
