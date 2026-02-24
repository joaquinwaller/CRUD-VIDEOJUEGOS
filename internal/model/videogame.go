package model

type Videogame struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Online bool   `json:"online"`
}
