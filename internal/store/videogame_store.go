package store

import (
	"CRUD-VIDEOJUEGOS/internal/model"
	"database/sql"
)

type Store interface {
	GetAll() ([]*model.Videogame, error)
	GetByID(id int) (*model.Videogame, error)
	Create(videogame *model.Videogame) (*model.Videogame, error)
	Update(id int, videogame *model.Videogame) (*model.Videogame, error)
	Delete(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

func (s *store) GetAll() ([]*model.Videogame, error) {
	q := `SELECT id, name, online FROM videogames`
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videogames []*model.Videogame

	for rows.Next() {
		v := &model.Videogame{}
		if err := rows.Scan(&v.ID, &v.Name, &v.Online); err != nil {
			return nil, err
		}
		videogames = append(videogames, v)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return videogames, nil
}

func (s *store) GetByID(id int) (*model.Videogame, error) {
	q := `SELECT id, name, online FROM videogames WHERE id = ?`

	v := &model.Videogame{}

	err := s.db.QueryRow(q, id).Scan(&v.ID, &v.Name, &v.Online)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return v, nil
}

func (s *store) Create(videogame *model.Videogame) (*model.Videogame, error) {
	q := `INSERT INTO videogames (name, online) VALUES (?, ?)`
	resp, err := s.db.Exec(q, videogame.Name, videogame.Online)
	if err != nil {
		return nil, err
	}
	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}
	videogame.ID = int(id)
	return videogame, nil

}

func (s *store) Update(id int, videogame *model.Videogame) (*model.Videogame, error) {
	q := `UPDATE videogames SET name = ?, online = ? WHERE id = ?`
	resp, err := s.db.Exec(q, videogame.Name, videogame.Online, id)
	if err != nil {
		return nil, err
	}
	affected, err := resp.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, nil
	}
	videogame.ID = id
	return videogame, nil
}

func (s *store) Delete(id int) error {
	q := `DELETE FROM videogames WHERE id = ?`
	resp, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}
	affected, err := resp.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
