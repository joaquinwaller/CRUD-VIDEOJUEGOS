package transport

import (
	"CRUD-VIDEOJUEGOS/internal/model"
	"CRUD-VIDEOJUEGOS/internal/service"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type VideogameHandler struct {
	service *service.Service
}

func New(s *service.Service) *VideogameHandler {
	return &VideogameHandler{
		service: s,
	}
}

// HandleVideogames maneja /videogames (lista y creación)
func (h *VideogameHandler) HandleVideogames(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		videogames, err := h.service.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(videogames); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		var v model.Videogame
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			http.Error(w, "error al parsear el body", http.StatusBadRequest)
			return
		}

		created, err := h.service.Create(&v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(created); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
	}
}

// HandleVideogamesByID maneja /videogames/{id}
func (h *VideogameHandler) HandleVideogamesByID(w http.ResponseWriter, r *http.Request) {
	// Ruta esperada: /videogames/{id}
	idStr := strings.TrimPrefix(r.URL.Path, "/videogames/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		v, err := h.service.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if v == nil {
			http.Error(w, "no encontrado", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(v); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		var body model.Videogame
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "error al parsear el body", http.StatusBadRequest)
			return
		}
		updated, err := h.service.Update(id, &body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if updated == nil {
			http.Error(w, "no encontrado", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(updated); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		if err := h.service.Delete(id); err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "no encontrado", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
	}
}