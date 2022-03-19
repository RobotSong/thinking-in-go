package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
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

	// Get user value 获取 URL 路径上的参数
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
	// 获取  ? 后的参数
	r.GET("/search", func(c *gin.Context) {
		keyword := c.Query("keyword")
		page := c.Query("page")
		fmt.Println("keyword:", keyword)
		fmt.Println("page:", page)
	})

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

// PostForm 获取 form 中的参数
func PostForm(r *gin.Engine) {
	r.POST("/api/form", func(c *gin.Context) {
		uname := c.PostForm("uname")
		pwd := c.PostForm("pwd")
		c.JSON(http.StatusOK, gin.H{
			"uname": uname,
			"pwd":   pwd,
		})
	})
}

type PostParams struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"mustBig"`
	Sex  int    `json:"sex" binding:"required"`
}

// PostJSON 从 请求体中 json 格式化成对象
func PostJSON(r *gin.Engine) {
	r.POST("/api/json", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "参数错误",
			})
			println(err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"form": p,
		})
	})
}

// PostFormSingleFile  单个文件上传
func PostFormSingleFile(r *gin.Engine) {
	r.POST("/api/single/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		category := c.PostForm("category")
		log.Println("分类:{}", category)
		// 上传文件至指定目录
		err := c.SaveUploadedFile(file, "D:\\tmp\\"+file.Filename)
		if err != nil {
			println("上传文件失败,", err.Error())
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
}

// PostFormMultipartFile  多个文件上传
func PostFormMultipartFile(r *gin.Engine) {
	r.POST("/api/multipart/upload", func(c *gin.Context) {
		// 单文件
		form, _ := c.MultipartForm()

		category := c.PostForm("category")
		log.Println("分类:", category)
		files := form.File["files"]

		fnames := ""
		for _, file := range files {
			// 上传文件至指定目录
			err := c.SaveUploadedFile(file, "D:\\tmp\\"+file.Filename)
			if err != nil {
				println("上传文件失败,", err.Error())
				continue
			}
			fnames += file.Filename + ";"
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", fnames))
	})
}

var mustBig validator.Func = func(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(int)
	if ok {
		if val < 18 {
			return false
		}
	}
	return true
}

func main() {
	gin.ForceConsoleColor()
	r := setupRouter()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mustBig", mustBig)
	}

	APIParam(r)
	PostForm(r)
	PostJSON(r)
	PostFormSingleFile(r)
	PostFormMultipartFile(r)
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8181")
}
