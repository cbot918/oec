package service

import (
	"oec/internal/repository"
	"oec/internal/types"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	R *repository.Repository
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		R: repository.NewRepository(db),
	}
}

func (s *Service) SignupService(user *types.SignupRequest) error {
	data := user
	// deal with jwt in the future

	return s.R.SignupRepository(data)
}

func (s *Service) InsertService() error {
	return s.R.InsertRepository()
}
