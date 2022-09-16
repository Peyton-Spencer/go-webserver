package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupAuthorized(r *gin.Engine) {
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		// AuthUserKey is an untyped constant (an interface)
		// .(string) is a type assertion
		// coerces any -> string
		// "any" is an alias for interface{}
		// := shorthand for var
		user := c.MustGet(gin.AuthUserKey).(string)
		// add "ok" bool to replace run-time panic
		// user, ok := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		// Bind the request body into the struct
		// nil value = no error
		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"value":  db[user],
			})
		}
	})
}
