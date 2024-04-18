package service

import "dev11/internal/repository"

type Service struct {
	Event
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Event: NewOrderService(repos.Event),
	}
}
