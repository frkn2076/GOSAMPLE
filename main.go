package main

import (
	_ "app/GoSample/config/environments"
	_ "app/GoSample/db"
	_ "app/GoSample/logger"
	_ "app/GoSample/infra/resource"
	"app/GoSample/router"
)

func main() {
	routing := router.SetupRouter()
	routing.Run(":5000")
}
