package bootstrap

import (
	"github.com/AliKhedmati/routate-backend/src/api"
	"github.com/AliKhedmati/routate-backend/src/cache"
	"github.com/AliKhedmati/routate-backend/src/config"
	"github.com/AliKhedmati/routate-backend/src/database"
)

func Init() {
	config.Init()
	cache.Init()
	database.Init()
	api.Init()
}
