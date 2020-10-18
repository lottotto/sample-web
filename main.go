package main

import (
	"math/rand"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	storeKey := rand.Int63n(10)
	store := cookie.NewStore([]byte(string(storeKey)))
	r.Use(sessions.Sessions("lottotto-sample-web-login-session", store))
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.POST("/login", login)
	r.GET("/logout", logout)
	r.GET("/github", loginGithub)
	r.GET("/auth/github/callback", callbackGithub)
	loginonly := r.Group("/secret")
	loginonly.Use(AuthRequired)
	{
		loginonly.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "OK"})
		})
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
