package repository

import "github.com/Venkat1abhinav/kairo/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
