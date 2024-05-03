package entityManager

import (
	"github.com/ZhanibekTau/go-jm-core/pkg/exception"
	"template_go/app/model"
)

func (m *Manager) GetCompanyTariffs() (*[]model.CompanyTariff, *exception.AppException) {
	return m.services.ICompanyTariffService.GetCompanyTariffs()
}
