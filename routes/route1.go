package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Envelope map[string]interface{}

func SampleRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/{name}", getSampleRoute)

	return router
}

func getSampleRoute(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	writeJSON(w, http.StatusOK, Envelope{"msg": "name provided " + name})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)

	if err != nil {
		return err
	}
	return nil
}
