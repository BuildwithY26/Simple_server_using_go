package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BuildwithY26/go/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}

func respondwithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Responding with 5XX error")
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondwithJSON(w, code, errorResponse{Error: msg})
}

func main() {
	feed, err := urlToFeed("https://wagslane.dev/index.xml")
	if err != nil {
		log.Fatal(err)
	}	
	fmt.Println(feed)
	
	godotenv.Load()

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	port := os.Getenv("PORT")

	conn, error := sql.Open("postgres", dbUrl)
	if error != nil {
		log.Fatal(error)
	}
	defer conn.Close()

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/ready", readinessHandler)
	v1Router.Get("/health", healthHandler)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.authMiddleware(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.authMiddleware(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.authMiddleware(apiCfg.handlerGetFeeds))
	v1Router.Post("/feed_follows", apiCfg.authMiddleware(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.authMiddleware(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.authMiddleware(apiCfg.handlerDeleteFeedFollow))
	router.Mount("/v1", v1Router)


srv := &http.Server{
	Handler: router,
	Addr:   ":" + port,
}

log.Printf("Server starting on port %s", port)
err = srv.ListenAndServe()
if err != nil {
	fmt.Printf("Failed to start server: %s\n", err)
}

	fmt.Printf("Port: %s\n", port)
}
