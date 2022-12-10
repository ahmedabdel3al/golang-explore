package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func pingMiddleWare(request *gin.Context) {
	var bVal bool
	bVal, _ = strconv.ParseBool(request.Query("useMiddleware"))
	
	if bVal == false {
		request.Next()
		return
	}
	request.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "unauthorized",
	})
}
func handleRouting() {
	router := gin.Default()
	router.GET("/ping", pingMiddleWare, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome to root system",
		})
	})
	router.GET("/home/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": id,
		})
	})
	router.Run()
}

func main() {
	handleRouting()
}
