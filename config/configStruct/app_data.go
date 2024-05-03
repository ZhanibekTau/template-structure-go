package structures

import (
	structures2 "github.com/ZhanibekTau/go-jm-core/pkg/config/structures"
	"github.com/ZhanibekTau/go-jm-core/pkg/modules/j/jStructures"
)

type AppData struct {
	RequestData  *jStructures.RequestData
	AppConfig    *structures2.AppConfig
	DbConfig     *structures2.DbConfig
	RedisConfig  *structures2.RedisConfig
	RestConfig   *RestConfig
	TokenConfig  *TokenConfig
	RabbitConfig *RabbitConfig
}
