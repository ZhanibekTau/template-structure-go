package model

import (
	"time"
)

type CompanyTariff struct {
	ID        int        `json:"id"`
	CompanyID int        `json:"company_id"`
	TariffID  int        `json:"tariff_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (c *CompanyTariff) TableName() string {
	return "company_tariffs"
}
