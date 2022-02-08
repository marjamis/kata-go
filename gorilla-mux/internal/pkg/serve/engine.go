package serve

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Serve is the entry of running the webserver that will use gorilla/mux
func Serve() {
	router := createMuxRouter()
	http.Handle("/", router)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:" + os.Getenv("PORT"),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	fmt.Printf("Server running on %s:%s\n", "127.0.0.1", os.Getenv("PORT"))
	log.Fatal(srv.ListenAndServe())
}

func createMuxRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products/{key}", ProductHandler)
	router.HandleFunc("/articles/{category}", ArticlesCategoryHandler)
	router.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	return router
}

// ProductHandler is used for the route that gets information about a product
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product: %v\n", vars["key"])
}

// ArticlesCategoryHandler is used for the route that gets information about an articles category
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

// ArticleHandler is used for the route that gets an article from a specific category
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v - Article: %v\n", vars["category"], vars["id"])
}
