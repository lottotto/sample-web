package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("lottotto-sample-web-login-session", store))

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.POST("/login", login)
	r.GET("/logout", logout)
	loginonly := r.Group("/secret")
	loginonly.Use(AuthRequired)
	{
		loginonly.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "OK"})
		})
	}

	return r
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username != "user" || password != "password" {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"status": "Authentication faild"})
		return
	}
	session := sessions.Default(c)
	session.Set("UserId", username)
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.HTML(200, "login.html", gin.H{"status": "Logout Complete!"})
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("UserId")
	if username == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

func main() {
	r := setupRouter()
	r.Run()
}
