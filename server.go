package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/deathwofl/wine-reviews/graph"
	"github.com/deathwofl/wine-reviews/graph/generated"
	"github.com/deathwofl/wine-reviews/pkg/domain"
	authmiddleware "github.com/deathwofl/wine-reviews/pkg/middleware"
	storage "github.com/deathwofl/wine-reviews/pkg/storage/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultPort = "8080"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")
	DBPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DBHost, DBUser, DBPassword, DBName, DBPort)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	userService := storage.UserService{DB: DB}

	// create router
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(authmiddleware.AuthMiddleware(userService))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// storage.MigrateModels(DB)
	d := domain.New(
		storage.WineService{DB: DB},
		storage.WineryService{DB: DB},
		userService,
		storage.ReviewService{DB: DB},
	)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Domain: d,
	}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "*"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
