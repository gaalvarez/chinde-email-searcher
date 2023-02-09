package mail

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func MailRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/search", func(w http.ResponseWriter, r *http.Request) {
		var searchParams SearchParams
		if err := json.NewDecoder(r.Body).Decode(&searchParams); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		result := QueryDataFromZincSearch(searchParams)
		json.NewEncoder(w).Encode(result)
	})

	return r
}
