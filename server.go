package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ravinder990011/product-management/graph"
	"github.com/ravinder990011/product-management/internal/db"
	"github.com/ravinder990011/product-management/internal/product"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()
	defer db.DB.Close()

	productRepo := product.NewRepository(db.DB)
	productService := product.NewService(productRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := mux.NewRouter()

	// GraphQL Playground
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// GraphQL Server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{ProductService: productService}}))
	r.Handle("/query", srv)

	// Start server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
