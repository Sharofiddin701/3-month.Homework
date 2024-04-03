package api

import (
	"rent-car/api/handler"
	"rent-car/service"
	"rent-car/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(services service.IServiceManager,store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store,services)

	r := gin.Default()

	r.POST("/car", h.CreateCar)
	r.GET("/car/:id", h.GetByIDCar)
	r.GET("/cars", h.GetAllCars)
	r.GET("/availablecars", h.GetAvaibleCars)
	r.PUT("/car/:id", h.UpdateCar)
	r.DELETE("/car/:id", h.DeleteCar)
	
	r.POST("/customer", h.CreateCustomer)
	r.GET("/customer/:id", h.GetByIDCustomer)
	r.GET("/customers", h.GetAllCustomer)
	r.PUT("/customer/:id", h.UpdateCustomer)
	r.DELETE("/customer/:id", h.DeleteCustomer)

	r.POST("/order", h.CreateOrder)
	r.GET("/order/:id", h.GetAllOrder)
	r.GET("/orders", h.GetAllOrder)
	r.PUT("/order/:id", h.UpdateOrder)
	r.DELETE("/order/:id", h.DeleteOrder)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
