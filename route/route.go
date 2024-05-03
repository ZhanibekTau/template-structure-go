package route

import (
	ginhelper "github.com/ZhanibekTau/go-jm-core/pkg/helpers/gin"
	"github.com/ZhanibekTau/go-jm-core/pkg/middleware"
	"github.com/ZhanibekTau/go-jm-core/pkg/modules/j/jMiddleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"template_go/app/handler/controller"
	structures "template_go/config/configStruct"
)

type Route struct {
	controller *controller.Controller
	appData    *structures.AppData
}

func NewRoute(controller *controller.Controller, appData *structures.AppData) *Route {
	return &Route{controller: controller, appData: appData}
}

func (r *Route) InitRoutes() *gin.Engine {
	router := ginhelper.InitRouter(r.appData.AppConfig)

	healthRouter := router.Group("/health")
	{
		healthRouter.Use(jMiddleware.ResponseHandler(r.appData.AppConfig))
		healthRouter.GET("alive", healthcheck)
		healthRouter.GET("ready", healthcheck)
	}

	sample := router.Group("sample/v1")
	{
		sample.Use(jMiddleware.RequestHandler(r.appData.RequestData))
		sample.Use(jMiddleware.ResponseHandler(r.appData.AppConfig))
		sample.Use(middleware.PinbaHandler(r.appData.AppConfig))
		sample.GET("/get-company-tariffs", r.controller.GetProducts)
	}

	return router
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "true", "status": http.StatusOK})
}
