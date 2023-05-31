package main

import (
	"basededatos/zinc"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	if len(os.Args) != 3 {
		log.Println("\x1b[31;1mIncorrect number of arguments: indexer <username> <password>\x1b[0m")
		return
	}

	log.Println("\x1b[32;1mSuccefully opened!! \x1b[0m")

	username := os.Args[1]
	password := os.Args[2]

	zinc.AuthValues(username, password)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5173"},
	})
	r.Use(corsOptions.Handler)

	// * Routes to search emails of the user
	r.Route("/search", func(r chi.Router) {
		// Subrouters:
		r.Route("/{index}/{searchTerm}/{from}", func(r chi.Router) {
			r.Get("/", searchHandler)
		})
	})

	http.ListenAndServe("192.168.1.7:3333", r)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	index := chi.URLParam(r, "index")
	term := chi.URLParam(r, "searchTerm")
	from := chi.URLParam(r, "from")
	email := zinc.Searcher(term, index, from)
	w.Write(email)
}