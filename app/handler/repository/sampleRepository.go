package repository

import (
	"github.com/ZhanibekTau/go-jm-core/pkg/exception"
	"github.com/jinzhu/gorm"
	"net/http"
	"template_go/app/model"
)

type CompanyTariffRepository struct {
	db *gorm.DB
}

func NewCompanyTariffRepository(db *gorm.DB) *CompanyTariffRepository {
	return &CompanyTariffRepository{db: db.LogMode(true)}
}

func (c *CompanyTariffRepository) GetCompanyTariffs() (*[]model.CompanyTariff, *exception.AppException) {
	var Products []model.CompanyTariff

	result := c.db.Find(&Products)
	if result.Error != nil {
		msg := result.Error
		return nil, exception.NewAppException(http.StatusBadRequest, msg)
	}

	return &Products, nil
}
