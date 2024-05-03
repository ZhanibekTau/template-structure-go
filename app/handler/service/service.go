package service

import (
	"github.com/ZhanibekTau/go-jm-core/pkg/exception"
	"template_go/app/handler/repository"
	"template_go/app/model"
)

type ICompanyTariffService interface {
	GetCompanyTariffs() (*[]model.CompanyTariff, *exception.AppException)
}

type Service struct {
	ICompanyTariffService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ICompanyTariffService: NewCompanyTariffService(repos.ICompanyTariffRepository),
	}
}
