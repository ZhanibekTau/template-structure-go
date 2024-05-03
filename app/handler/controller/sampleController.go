package controller

import (
	"github.com/ZhanibekTau/go-jm-core/pkg/exception"
	httpResponse "github.com/ZhanibekTau/go-jm-core/pkg/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ct *Controller) GetProducts(c *gin.Context) {
	companyTariffs, err := ct.manager.GetCompanyTariffs()
	if err != nil {
		httpResponse.Error(c, exception.NewAppException(http.StatusInternalServerError, err.Error))
		return
	}

	httpResponse.Response(c, http.StatusOK, companyTariffs)
}
