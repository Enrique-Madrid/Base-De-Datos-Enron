package backend

import (
	"basededatos/zinc"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	log.Println("\x1b[32;1mBackend successfully opened!! \x1b[0m")

	// * Setting the environment variables

	env := os.Environ()
	for _, value := range env {
		if strings.HasPrefix(value, "ZINC_") {
			split := strings.Split(value, "=")
			if len(split) != 2 {
				log.Println("Error parsing environment variable: ", value)
				return
			}
			os.Setenv(split[0], split[1])
		}
	}
	username := os.Getenv("ZINC_FIRST_ADMIN_USER")
	password := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")

	if username == "" || password == "" {
		log.Println("Error: Missing environment variables")
		return
	}

	zinc.AuthValues(username, password)

	// * Setting the server

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		Debug:          true,
	})
	r.Use(corsOptions.Handler)

	// * Routes to search emails of the user
	r.Route("/search", func(r chi.Router) {
		// Subrouters:
		r.Route("/{searchTerm}/{from}", func(r chi.Router) {
			r.Get("/", searchHandler)
		})
	})
	err := http.ListenAndServe("192.168.1.7:3333", r)
	if err != nil {
		log.Println("\x1b[31;1mError starting server: \x1b[0m", err)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	term := chi.URLParam(r, "searchTerm")
	from := chi.URLParam(r, "from")
	email := zinc.Searcher(term, from)

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(email)
	if err != nil {
		log.Println("\x1b[31;1mError writing response: \x1b[0m", err)
	}
}

func Start() {
	main()
}
