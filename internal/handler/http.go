package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"gladiatortravel/internal/model"
	"gladiatortravel/internal/usecase"
)

type HTTPHandler struct {
	svc *usecase.Service
}

func NewHTTPHandler(svc *usecase.Service) *HTTPHandler {
	return &HTTPHandler{svc: svc}
}

func (h *HTTPHandler) Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/health", h.health)
	r.Get("/swagger", h.swaggerUI)
	r.Get("/swagger/", h.swaggerUI)
	r.Get("/swagger/openapi.yaml", h.swaggerSpec)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/destinations", h.listDestinations)
		r.Get("/destinations/{id}/venues", h.listVenues)
		r.Get("/destinations/{id}/accommodations", h.listAccommodations)
		r.Get("/destinations/{id}/travel-options", h.listTravelOptions)
		r.Put("/me/preferences", h.upsertPreferences)
		r.Post("/me/venues/{venueID}/feedback", h.venueFeedback)
		r.Post("/trips", h.createTrip)
		r.Post("/trips/{tripID}/generate-itinerary", h.generateItinerary)
		r.Get("/trips/{tripID}/itinerary", h.getItinerary)
	})
	return r
}

func (h *HTTPHandler) swaggerUI(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>GladiatorTravel Swagger</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script>
      window.ui = SwaggerUIBundle({
        url: "/swagger/openapi.yaml",
        dom_id: "#swagger-ui"
      });
    </script>
  </body>
</html>`
	_, _ = fmt.Fprint(w, html)
}

func (h *HTTPHandler) swaggerSpec(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "api/openapi.yaml")
}

func (h *HTTPHandler) health(w http.ResponseWriter, _ *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := h.svc.Health(ctx); err != nil {
		respondError(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *HTTPHandler) listDestinations(w http.ResponseWriter, r *http.Request) {
	sortBy := r.URL.Query().Get("sort")
	minScore := parseFloatDefault(r.URL.Query().Get("min_score"), 0)
	items, err := h.svc.ListDestinations(r.Context(), sortBy, minScore)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, items)
}

func (h *HTTPHandler) listVenues(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid destination id")
		return
	}
	venueType := r.URL.Query().Get("type")
	var openLate *bool
	if raw := r.URL.Query().Get("open_late"); raw != "" {
		v := raw == "true"
		openLate = &v
	}
	var priceTier *int
	if raw := r.URL.Query().Get("price_tier"); raw != "" {
		v, err := strconv.Atoi(raw)
		if err == nil {
			priceTier = &v
		}
	}
	items, err := h.svc.ListVenues(r.Context(), id, venueType, openLate, priceTier)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, items)
}

func (h *HTTPHandler) listAccommodations(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid destination id")
		return
	}
	mode := r.URL.Query().Get("mode")
	items, err := h.svc.ListAccommodations(r.Context(), id, mode)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, items)
}

func (h *HTTPHandler) listTravelOptions(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid destination id")
		return
	}
	origin := r.URL.Query().Get("origin_city")
	items, err := h.svc.ListTravelOptions(r.Context(), id, origin)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, items)
}

func (h *HTTPHandler) upsertPreferences(w http.ResponseWriter, r *http.Request) {
	var p model.UserPreferences
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		respondError(w, http.StatusBadRequest, "invalid body")
		return
	}
	if p.UserID == 0 {
		respondError(w, http.StatusBadRequest, "user_id is required")
		return
	}
	if err := h.svc.UpsertPreferences(r.Context(), p); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"status": "saved"})
}

func (h *HTTPHandler) venueFeedback(w http.ResponseWriter, r *http.Request) {
	venueID, err := strconv.ParseInt(chi.URLParam(r, "venueID"), 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid venue id")
		return
	}
	var req struct {
		UserID   int64  `json:"user_id"`
		Feedback string `json:"feedback"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid body")
		return
	}
	if err := h.svc.AddVenueFeedback(r.Context(), req.UserID, venueID, req.Feedback); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, map[string]string{"status": "recorded"})
}

func (h *HTTPHandler) createTrip(w http.ResponseWriter, r *http.Request) {
	var t model.Trip
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		respondError(w, http.StatusBadRequest, "invalid body")
		return
	}
	id, err := h.svc.CreateTrip(r.Context(), t)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, map[string]int64{"trip_id": id})
}

func (h *HTTPHandler) generateItinerary(w http.ResponseWriter, r *http.Request) {
	tripID, err := strconv.ParseInt(chi.URLParam(r, "tripID"), 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid trip id")
		return
	}
	var req struct {
		DestinationID int64 `json:"destination_id"`
		Days          int   `json:"days"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid body")
		return
	}
	if err := h.svc.GenerateItinerary(r.Context(), tripID, req.DestinationID, req.Days); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, map[string]string{"status": "generated"})
}

func (h *HTTPHandler) getItinerary(w http.ResponseWriter, r *http.Request) {
	tripID, err := strconv.ParseInt(chi.URLParam(r, "tripID"), 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid trip id")
		return
	}
	items, err := h.svc.GetItinerary(r.Context(), tripID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, items)
}

func parseFloatDefault(raw string, fallback float64) float64 {
	if raw == "" {
		return fallback
	}
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return fallback
	}
	return v
}

func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func respondError(w http.ResponseWriter, status int, msg string) {
	respondJSON(w, status, map[string]string{"error": msg})
}
