package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true}))
	r.POST("/user", func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	err := r.Run()
	if err != nil {
		return
	}

	r.Run("178.170.194.41:8080")
}
