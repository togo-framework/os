package os

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Service) mountRoutes(r chi.Router) {
	r.Route("/api/os", func(r chi.Router) {
		r.Get("/session", func(w http.ResponseWriter, req *http.Request) {
			writeJSON(w, http.StatusOK, s.Prefs(req.Context(), currentUser(req)))
		})
		r.Put("/session", func(w http.ResponseWriter, req *http.Request) {
			var p DesktopPrefs
			if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			err := s.SavePrefs(req.Context(), currentUser(req), p)
			writeErr(w, err, p)
		})
		r.Get("/apps", func(w http.ResponseWriter, req *http.Request) {
			writeJSON(w, http.StatusOK, Apps(s.k))
		})
		r.Get("/apps/{slug}", func(w http.ResponseWriter, req *http.Request) {
			a, ok := AppBySlug(s.k, chi.URLParam(req, "slug"))
			if !ok {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			writeJSON(w, http.StatusOK, a)
		})
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, err error, ok any) {
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, ok)
}
