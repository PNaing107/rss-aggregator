package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	// configure cors
	router.Use(cors.Handler(Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.HandleFunc("/healthz", handlerReadiness)

	// nest the v1 router under the /v1 path
	router.Mount("/v1", v1Router)


	server := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("server starting on port %v", portString)

	err := server.ListenAndServe() // this is a blocking method

	if err != nil {
		log.Fatal(err)
	}

}