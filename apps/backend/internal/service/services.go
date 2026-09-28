package service

import (
	"github.com/Venkat1abhinav/kairo/internal/lib/job"
	"github.com/Venkat1abhinav/kairo/internal/repository"
	"github.com/Venkat1abhinav/kairo/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
