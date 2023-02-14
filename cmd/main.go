package main

import (
	"context"
	"inventory_graphql_api/config"
	"inventory_graphql_api/graph"
	"inventory_graphql_api/internal/repository"
	"inventory_graphql_api/postgres"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx := context.Background()

	conf := config.New()

	db := postgres.New(ctx, conf)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Timeout(time.Second * 60))

	// repository
	barangMasukRepo := repository.NewBarangMasukRepo(db)
	barangKeluarRepo := repository.NewBarangKeluarRepo(db)

	r.Handle("/", playground.Handler("Inventory barang Graphql API", "/query"))
	r.Handle("/query", handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					BarangMasukRepo:  barangMasukRepo,
					BarangKeluarRepo: barangKeluarRepo,
				},
			},
		),
	))

	log.Fatal(http.ListenAndServe(":3333", r))
	log.Println("server running")
}
