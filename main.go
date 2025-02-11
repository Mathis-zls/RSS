package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Mathis-zls/RSS/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portSring := os.Getenv("PORT")
	if portSring == "" {
		log.Fatal("Port missing")
	}

	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("DB missing")
	}

	conn, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("Error Connecting to Database")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healtz", handlerReady)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portSring,
	}
	log.Printf("Server start on port %v", portSring)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal()
	}

}
