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
	router.HandleFunc("/products/{key}", ProductHandler).
		Methods("GET").
		Schemes("http")

	// Creates a sub route which can be used to "namespace" matching routes
	// In this case I've created ones for "/articles/" and then adding routes to the subrouter
	articleSubRouter := router.PathPrefix("/articles/").Subrouter()
	articleSubRouter.HandleFunc("/{category}", ArticlesCategoryHandler)
	articleSubRouter.HandleFunc("/{category}/{id:[0-9]+}", ArticleHandler)

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
