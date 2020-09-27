package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"useragent": c.GetHeader("User-Agent"),
		})
	})
	r.StaticFile("/favicon.ico", "./favicon.ico")
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
