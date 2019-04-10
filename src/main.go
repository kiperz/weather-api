package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
	"github.com/kiperz/weather-api/client"
	"github.com/kiperz/weather-api/endpoints"
	"github.com/kiperz/weather-api/repositories"
	"log"
	"net/http"
	"os"
)

func SetupRouter(bookmarkRepository *repositories.PlaceBookmarkRepository, queryRepository *repositories.PlaceQueryRepository, weatherClient *client.Client) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,                             // Log API request calls
		middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
		middleware.Recoverer,                          // Recover from panics without crashing server
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/bookmark", endpoints.BookmarksRoutes(bookmarkRepository))
		r.Mount("/api/query", endpoints.QueriesRoutes(queryRepository, weatherClient))
	})

	return router
}

var (
	dbHost string
	dbName string
	dbUser string
	dbPassword string
	apiKey string
)

func init() {
	var found bool
	dbHost, found = os.LookupEnv("DB_HOST")
	if !found {
		dbHost = "localhost"
	}

	dbName, found = os.LookupEnv("DB_NAME")
	if !found {
		dbName = "weatherapi"
	}

	dbUser, found = os.LookupEnv("DB_USER")
	if !found {
		dbUser = "postgres"
	}

	dbPassword, found = os.LookupEnv("DB_PASSWORD")
	if !found {
		dbPassword = "postgres"
	}

	apiKey, found = os.LookupEnv("API_KEY")
	if !found {
		panic("No api key set!")
		return
	}
}

func main() {
	var (
		placeBookmarkRepository *repositories.PlaceBookmarkRepository
		placeQueryRepository *repositories.PlaceQueryRepository
		weatherClient *client.Client
		db *sql.DB
		err error
	)
	weatherClient = client.New(apiKey)
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbName, dbUser, dbPassword)
	if db, err = sql.Open("postgres", connStr); err != nil {
		fmt.Sprintln("Error ocurred while connecting to database: %s", err.Error());
		return
	}

	if placeBookmarkRepository, err = repositories.NewPlaceBookmarkRepository(db); err != nil {
		fmt.Printf("Error ocurred while creating PlaceBookmarkRepository: %s\n", err)
		return
	}


	if placeQueryRepository, err = repositories.NewPlaceQueryRepository(db); err != nil {
		fmt.Printf("Error ocurred while creating PlaceQueryRepository: %s\n", err)
		return
	}
	router := SetupRouter(placeBookmarkRepository, placeQueryRepository, weatherClient)
	log.Fatal(http.ListenAndServe(":8080", router))
}