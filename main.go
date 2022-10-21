package main

import (
	"back_end_for_pizza/configs"

	"back_end_for_pizza/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	configs.ConnectDB()

	routes.UserRouter(router)

	router.Run("127.0.0.1:8080")

}
