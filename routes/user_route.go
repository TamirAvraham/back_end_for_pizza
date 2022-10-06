package routes

import (
	"github.com/gin-gonic/gin"

	"back_end_for_pizza/controllers"
)

func UserRouter(router *gin.Engine) {
	router.POST("/NewPizza", controllers.CreatePizza())
	router.GET("/GetPizza/:id", controllers.GetPizza())
	router.PUT("/UpdatePizza/:id", controllers.EditPizza())
	router.DELETE("/DeletePizza/:id", controllers.DeletePizza())
	router.GET("/GetAllPizzas", controllers.GetAllPizzas())
	router.POST("/NewOrder/:pizzaId", controllers.CreateOrder())
	router.POST("/NewOrderWithNewPizza", controllers.CreateOrderWithNewPizza())
	router.GET("/GetOrders", controllers.GetAllOrders())
	router.DELETE("/DeleteOrder/:id", controllers.DeleteOrder())
}
