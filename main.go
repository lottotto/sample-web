package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"useragent": c.GetHeader("User-Agent"),
		})
	})
	r.GET("/list", getAllEmployee)
	r.POST("/login", login)
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")
	r.Static("/assets", "./assets")
	return r
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username != "user" || password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication faild"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func main() {
	r := setupRouter()
	r.Run()
}
