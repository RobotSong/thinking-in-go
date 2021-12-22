package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

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
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	r.PUT("/myput")

	return r
}

// APIParam 获取 URL 中的API 参数
func APIParam(r *gin.Engine) {
	r.GET("/api/:name/:action/*value", func(c *gin.Context) {
		// API 中，可以使用 :xxx *xxx ，但是只能最后一个是 *xxx
		// /api/zhangsan/abc
		// :name 会直接获取到  zhangsan 的值
		// *value 会带上 / 得到 /abc
		name := c.Param("name")
		val := c.Param("value")
		action := c.Param("action")
		result := fmt.Sprintf("%s is %s action is %s", name, val, action)
		c.String(http.StatusOK, result)
	})
}

func main() {
	gin.ForceConsoleColor()
	r := setupRouter()
	APIParam(r)
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}
