package entityManager

import (
	"template_go/app/handler/service"
)

type Manager struct {
	services *service.Service
}

func NewManager(services *service.Service) *Manager {
	return &Manager{
		services: services,
	}
}
