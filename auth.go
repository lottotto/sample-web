package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

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
