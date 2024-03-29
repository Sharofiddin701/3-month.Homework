package api

import (
	"rent-car/api/handler"
	"rent-car/storage"
	"rent-car/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

func New(services service.IServiceManager, store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store, services)

	r := gin.Default()

	r.POST("/car", h.CreateCar)
	// r.GET("/car/:id", h.GetAllCars)
	r.GET("/car", h.GetAllCars)
	r.GET("/car/:id", h.GetByIDCar)
	r.GET("/avacar", h.GetAvailableCars)
	r.PUT("/car/:id", h.UpdateCar)
	r.DELETE("/car/:id", h.DeleteCar)
	// r.PATCH("/car/:id", h.UpdateUserPassword)

	r.GET("/customer/:id", h.GetByIDCustomer)
	r.GET("/customer", h.GetAllCustomer)
	r.POST("/customer", h.CreateCustomer)
	r.PUT("/customer/:id", h.UpdateCustomer)
	r.DELETE("/customer/:id", h.DeleteCustomer)

	r.GET("/order", h.GetAllOrder)
	r.GET("/order/:id", h.GetOne)
	r.POST("/order", h.CreateOrder)
	r.PUT("/order/:id", h.UpdateOrder)

	// r.GET("/getcus/:id", h.GetCusCars)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
