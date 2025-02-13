package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

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
	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}
	go startScraping(db, 10, time.Minute)

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
	v1Router.Get("/users", apiCfg.middelwareAuth(apiCfg.handlerGetUserByApiKey))

	v1Router.Post("/feeds", apiCfg.middelwareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	v1Router.Post("/feedfollow", apiCfg.middelwareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feedfollow", apiCfg.middelwareAuth(apiCfg.handlerGetFeedsFollows))
	v1Router.Delete("/feedfollow/{feedFollowID}", apiCfg.middelwareAuth(apiCfg.handlerDeleteFeedFollow))

	v1Router.Get("/posts", apiCfg.middelwareAuth(apiCfg.handlerGetPostForUser))

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
