package service

import (
	"github.com/ZhanibekTau/go-jm-core/pkg/exception"
	"template_go/app/handler/repository"
	"template_go/app/model"
)

type CompanyTariffService struct {
	repo repository.ICompanyTariffRepository
}

func NewCompanyTariffService(repo repository.ICompanyTariffRepository) *CompanyTariffService {
	return &CompanyTariffService{repo: repo}
}

func (c *CompanyTariffService) GetCompanyTariffs() (*[]model.CompanyTariff, *exception.AppException) {
	return c.repo.GetCompanyTariffs()
}
