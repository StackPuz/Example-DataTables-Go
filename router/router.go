package router

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	productController := controllers.ProductController{}
	router := gin.Default()
	router.StaticFile("/", "./public/index.html")
	router.GET("/api/products", productController.Index)
	router.Run()
}
