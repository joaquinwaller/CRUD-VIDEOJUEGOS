package service

import (
	"CRUD-VIDEOJUEGOS/internal/model"
	"CRUD-VIDEOJUEGOS/internal/store"
	"errors"
)

type Logger interface {
	Log(message string)
}

type Service struct {
	store  store.Store
	logger Logger
}

func New(st store.Store) *Service {
	return &Service{
		store:  st,
		logger: nil,
	}
}

func (s *Service) GetAll() ([]*model.Videogame, error) {
	if s.logger != nil {
		s.logger.Log("Estamos obteniendo los juegos")
	}
	return s.store.GetAll()
}

func (s *Service) GetByID(id int) (*model.Videogame, error) {
	return s.store.GetByID(id)
}

func (s *Service) Create(videogame *model.Videogame) (*model.Videogame, error) {
	if videogame.Name == "" {
		return nil, errors.New("Necesitamos Nombre")

	}

	return s.store.Create(videogame)
}

func (s *Service) Update(id int, videogame *model.Videogame) (*model.Videogame, error) {
	if videogame.Name == "" {
		return nil, errors.New("Necesitamos Nombre")

	}
	return s.store.Update(id, videogame)
}

func (s *Service) Delete(id int) error {
	return s.store.Delete(id)
}
