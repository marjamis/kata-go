package example

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
	go http.ListenAndServe(":8000", nil)

	seconds := time.Duration(20)
	fmt.Printf("Webserver open on port localhost:8000. And will be continue to run for: %d seconds\n", seconds)

	for i := seconds; i >= 0; i-- {
		fmt.Printf("%d second/s left...\n", i)
		time.Sleep(time.Second)
	}
	fmt.Println("Webserver stopping...")
}

func init() {
	category := GetCategories().AddCategory("webserver")

	category.AddExample("basic",
		CategoryExample{
			Description: "Basic Webserver",
			Function:    runWebserver,
		})
}
