package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	// example: authorized endpoints
	setupAuthorized(r)
	r.GET("/user/:name", getUserValue)

	// example: websocket
	r.LoadHTMLFiles("index.html")
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
	r.GET("/ws", echoHandler)

	return r
}

var getUserValue gin.HandlerFunc = func(c *gin.Context) {
	// :name in the URL is the Param
	user := c.Params.ByName("name")

	value, ok := db[user]
	if ok {
		// JSON returns in alphabetical order
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}
