package controller

import (
	"template_go/app/handler/entityManager"
	structures "template_go/config/configStruct"
)

type Controller struct {
	manager *entityManager.Manager
	appData *structures.AppData
}

func NewController(manager *entityManager.Manager, appData *structures.AppData) *Controller {
	return &Controller{manager: manager, appData: appData}
}
