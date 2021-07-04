package main

import (
	"app/GoSample/config/environments"
	"app/GoSample/config/session"
	"app/GoSample/logger"
	"app/GoSample/db"
	"app/GoSample/infra/resource"
	"app/GoSample/router"
)

func main() {
	environments.Load()
	session.Start()
	logger.Start()

	db.ConnectDatabases()
	db.MigrateTables()
	db.InitScripts()

	resource.CacheFeeder()

	routing := router.SetupRouter()
	routing.Run(":5000")
}
