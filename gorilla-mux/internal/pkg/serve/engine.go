package serve

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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

	log.Printf("Server running on %s:%s\n", "127.0.0.1", os.Getenv("PORT"))
	log.Fatal(srv.ListenAndServe())
}

func createMuxRouter() *mux.Router {
	// For more information about the matched, subrouters, etc. be sure to visit
	// the official documentation: https://github.com/gorilla/mux

	// Creates a new parent router
	router := mux.NewRouter()

	// Creates a new router on the parent router that will use the ProductHandler for processing
	// if the Methods and Schemes match what's defined
	router.HandleFunc("/products/{key}", ProductHandler).
		Methods("GET").
		Schemes("http")
	// This logging middleware will be used for all routes configured under this specific router but
	// not subrouters
	router.Use(loggingMiddleware)

	// Creates a subrouter off of the parent which can be used to "namespace" matching routes
	// In this case I've created ones for "/articles/" and are then adding subroutes to the "/articles" path
	articlesSubRouter := router.PathPrefix("/articles/").Subrouter()
	articlesSubRouter.HandleFunc("/{category}", ArticlesCategoryHandler)
	articlesSubRouter.HandleFunc("/{category}/{id:[0-9]+}", ArticleHandler)
	// This middleware will be used for routes under the articlesSubRouter on top off any other
	// middleware that may be called before it
	articlesSubRouter.Use(blockingMiddleware)

	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func blockingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is a simple check which will not call next.ServerHTTP(), hence not execute
		// the normal HandleFunc for the path, essentially stopping the normal processing of the request
		magicNumber := "6"
		if strings.Contains(r.URL.Path, magicNumber) {
			log.Printf("The requests path contains the number %s, Aborting request...\n", magicNumber)
			// I'm writing to the ResponseWriter to generate the InternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
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
