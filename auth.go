package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"
)

var githubOauthConf = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/github/callback",
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	Scopes:       []string{"user"},
	Endpoint:     oauth2github.Endpoint,
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

// AuthRequired ...
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("UserId")
	if username == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

func loginGithub(c *gin.Context) {

	// b := make([]byte, 16)
	// rand.Read(b)
	// state := base64.URLEncoding.EncodeToString(b)

	c.SetCookie("oauthState", "state1234", 1000000000, "/", "localhost:8080", false, false)

	authURL := githubOauthConf.AuthCodeURL("state1234")
	fmt.Println("authURL: " + authURL)
	c.Redirect(http.StatusMovedPermanently, authURL)
}

func callbackGithub(c *gin.Context) {
	oauthState, err := c.Cookie("oauthState")
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	if oauthState != c.Query("state") {
		log.Println("invalid oauth github state")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Query("code")
	token, err := githubOauthConf.Exchange(context.Background(), code)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	ts := oauth2.StaticTokenSource(token)
	tc := oauth2.NewClient(c, ts)
	client := github.NewClient(tc)
	user, _, err := client.Users.Get(c, "")
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	c.JSON(200, gin.H{"username": user.Name, "userID": user.ID})
}
