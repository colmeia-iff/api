package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Starting")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return
	}

	r := chi.NewRouter()
	m := chiprometheus.NewMiddleware("router")

	r.Use(m)
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})
	r.Use(corsOptions.Handler)
	r.Mount("/user", UserRouter())
	r.Mount("/apiary", ApiaryRouter())
	r.Mount("/hive", HiveRouter())
	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}
