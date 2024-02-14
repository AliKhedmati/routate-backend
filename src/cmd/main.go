package main

import (
	"github.com/AliKhedmati/routate-backend/config"
	"github.com/AliKhedmati/routate-backend/src/api"
	"github.com/AliKhedmati/routate-backend/src/database"
)

func main() {
	config.Init()
	database.Init()
	api.Init()

	//collection := database.GetDB()
}
