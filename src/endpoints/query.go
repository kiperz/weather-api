package endpoints


import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/kiperz/weather-api/client"
	"github.com/kiperz/weather-api/repositories"
	"net/http"
	"time"
)

type QueryRequestBody struct {
	Name string
}

func QueriesRoutes(repository *repositories.PlaceQueryRepository, weatherClient *client.Client) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/stats", GetStatistics(repository))
	router.Post("/", CreateQuery(repository, weatherClient))
	return router
}

func GetStatistics(repository *repositories.PlaceQueryRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			sendError("missing name query parameter", w, r)
			return
		}
		stats, err := repository.GetStatistics(name)
		if err != nil {
			sendError(fmt.Sprintf(`error getting bookmark data: %d`, err.Error()), w, r)
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		render.JSON(w, r, stats)
	}
}

func CreateQuery(repository *repositories.PlaceQueryRepository, weatherClient *client.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body *QueryRequestBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendError(fmt.Sprintf(`error while json parse: %d`, err.Error()), w, r)
		}
		data, err := weatherClient.GetWeatherData(body.Name)
		if err != nil {
			sendError(fmt.Sprintf(`error while getting data from api: %d`, err.Error()), w, r)
		}
		query := &repositories.PlaceQuery{
			Name: body.Name,
			Type: data.Type,
			Temperature: data.Temperature,
			MinimumTemperature: data.MinimumTemperature,
			MaximumTemperature: data.MaximumTemperature,
			Date: time.Now().Format("2006-01-02"),
		}
		if err := repository.Put(query); err != nil {
			sendError(fmt.Sprintf(`error while saving data: %d`, err.Error()), w, r)
		}
		w.WriteHeader(200)
	}
}