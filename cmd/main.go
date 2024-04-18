package main

import (
	"fmt"
	"log"
	"net/http"

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

	port := "9999"

	fmt.Println("lissen port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
