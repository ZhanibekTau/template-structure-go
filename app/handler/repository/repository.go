package repository

import (
	"github.com/ZhanibekTau/go-jm-core/pkg/exception"
	"github.com/jinzhu/gorm"
	"template_go/app/model"
)

type ICompanyTariffRepository interface {
	GetCompanyTariffs() (*[]model.CompanyTariff, *exception.AppException)
}

type Repository struct {
	ICompanyTariffRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		ICompanyTariffRepository: NewCompanyTariffRepository(db),
	}
}
