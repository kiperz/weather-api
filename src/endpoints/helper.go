package endpoints

import (
	"github.com/go-chi/render"
	"net/http"
)

func sendError(message string, w http.ResponseWriter, r *http.Request) {
	error := make(map[string]string)
	error["status"] = "error"
	error["message"] = message
	w.WriteHeader(500)
	render.JSON(w, r, error)
}
