package server

import (
	"github.com/gin-gonic/gin"
	"github.com/koko2824/first_postgresql/controller"
)

func Init()  {
	r := router()
	r.Run("localhost:8080")
}

func router() *gin.Engine {
	r := gin.Default()
	ctrl := controller.Controller{}

	// GET
	r.GET("/users", ctrl.GetAll)
	r.GET("/user/:id", ctrl.GetOne)

	// POST
	r.POST("/user", ctrl.Insert)

	// PUT
	r.PUT("/user/:id", ctrl.Update)

	// DELETE
	r.DELETE("/user/:id", ctrl.Delete)
	return r
}