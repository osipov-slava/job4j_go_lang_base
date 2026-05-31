package api

import "job4j.ru/go-lang-base/internal/tracker/repository"

type Server struct {
	Repository *repository.RepoPg
}

func NewServer(repo *repository.RepoPg) *Server {
	return &Server{Repository: repo}
}
