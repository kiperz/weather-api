package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/kiperz/weather-api/repositories"
	"net/http"
)

type CreateBookmarkRequestBody struct {
	Name string
}

func BookmarksRoutes(repository *repositories.PlaceBookmarkRepository) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetBookmark(repository))
	router.Post("/", CreateBookmark(repository))
	return router
}

func GetBookmark(repository *repositories.PlaceBookmarkRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bookmarks, err := repository.Get()
		if err != nil {
			sendError(fmt.Sprintf(`error getting bookmark data: %d`, err.Error()), w, r)
		}
		w.WriteHeader(200)
		render.JSON(w, r, bookmarks)
	}
}

func CreateBookmark(repository *repositories.PlaceBookmarkRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body *CreateBookmarkRequestBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendError(fmt.Sprintf(`error while json parse: %d`, err.Error()), w, r)
		}
		if err := repository.Put(&repositories.PlaceBookmark{
			Name: body.Name,
		}); err != nil {
			sendError(fmt.Sprintf(`error while saving data: %d`, err.Error()), w, r)
		}
		w.WriteHeader(200)
	}
}