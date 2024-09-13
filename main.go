// main.go
package main

import (
	"example.com/gonicginrestapi/routers"
	"example.com/gonicginrestapi/utils"
)

func main() {
	utils.LogInfo("Main Begin")

	router := routers.InitRoutes()

	router.Run(":3002")
}
