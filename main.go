package main

import (
	"context"
	"github.com/ZhanibekTau/go-jm-core/pkg/helpers/db/mysql"
	"github.com/ZhanibekTau/go-jm-core/pkg/modules/j/jStructures"
	"template_go/app/handler/controller"
	"template_go/app/handler/entityManager"
	"template_go/app/handler/repository"
	"template_go/app/handler/service"
	"template_go/config"
	structures "template_go/config/configStruct"
	"template_go/route"
	"template_go/server"
	"time"

	_ "time/tzdata"

	"github.com/sirupsen/logrus"
)

func init() {
	location, err := time.LoadLocation("Asia/Almaty")
	if err != nil {
		logrus.Fatalf("error loading location: %v", err.Error())
	}

	time.Local = location
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	appConfig, dbConfig, redisConfig, restConfig, tokenConfig, rabbitConfig, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("error database connection: %v", err.Error())
	}

	appData := &structures.AppData{
		AppConfig:   appConfig,
		DbConfig:    dbConfig,
		RedisConfig: redisConfig,
		RestConfig:  restConfig,
		RequestData: &jStructures.RequestData{
			ServiceName: appConfig.Name,
		},
		TokenConfig:  tokenConfig,
		RabbitConfig: rabbitConfig,
	}

	db, err := mysql.NewGormSqlDB(dbConfig)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer mysql.CloseDbConnection(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	manager := entityManager.NewManager(services)
	controllers := controller.NewController(manager, appData)
	routes := route.NewRoute(controllers, appData)

	srv := new(server.Server)

	if err := srv.Run(appData.AppConfig.ServerAddress, routes.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
