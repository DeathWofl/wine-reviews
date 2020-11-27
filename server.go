package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/deathwofl/wine-reviews/graph"
	"github.com/deathwofl/wine-reviews/graph/generated"
	storage "github.com/deathwofl/wine-reviews/pkg/storage/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultPort = "8080"

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	dsn := "host=localhost user=postgres password=mysecretpassword dbname=winereviews port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// storage.MigrateModels(db)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UserService:   storage.UserService{DB: DB},
		WineryService: storage.WineryService{DB: DB},
		WineService:   storage.WineService{DB: DB},
		ReviewService: storage.ReviewService{DB: DB},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
