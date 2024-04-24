package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Starting")

	r := chi.NewRouter()
	m := chiprometheus.NewMiddleware("router")

	r.Use(m)
	corsOptions := cors.AllowAll()
	r.Use(corsOptions.Handler)
	r.Mount("/user", UserRouter())
	r.Mount("/apiary", ApiaryRouter())
	r.Mount("/hive", HiveRouter())

	// Use envPortOr function to determine the port
	port := envPortOr("9999")

	fmt.Println("listen port: ", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// envPortOr returns PORT from environment if found, defaults to
// value in `port` parameter otherwise.
func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		fmt.Println(envPort)
		return ":" + envPort
	}
	return ":" + port
}
