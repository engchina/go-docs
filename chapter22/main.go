package main

import (
	"chapter22/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func myHandler(c *gin.Context) {
	fmt.Println("myHandler")
}

func myBaseHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.FullPath()
		method := ctx.Request.Method
		fmt.Println("myBaseHandler", path, method)
	}
}

func someGet(c *gin.Context) {
	fmt.Println("someGet")
	_, err := c.Writer.WriteString("someGet")
	if err != nil {
		return
	}
}

func somePost(c *gin.Context) {
	fmt.Println("somePost")
	_, err := c.Writer.WriteString("somePost")
	if err != nil {
		return
	}
}

func getKey(c *gin.Context) {
	name := c.DefaultQuery("name", "world")
	c.String(http.StatusOK, "%s %s", "hello", name)
}

func postKey(c *gin.Context) {
	name := c.PostForm("name")
	age := c.DefaultPostForm("age", "18")
	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"name": name,
		"age":  age,
	})
	//c.String(http.StatusOK, "%s %s", "hello", name)
}

func getParam(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "%s %s", "hello", name)
}

func goRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "goRegister.html", gin.H{})
}

func register(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		return
	}
	c.String(http.StatusOK, "%s", user)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	r := gin.Default()
	r.Use(myHandler, myBaseHandler())
	r.Static("../static", "static")
	//r.Static("../static/favicon.ico", "static/favicon.ico")
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "Pong"})
	})

	r.GET("/someGet", someGet)
	r.POST("/somePost", somePost)
	r.GET("/getKey", getKey)
	r.POST("/postKey", postKey)
	r.GET("/getParam/:name", getParam)
	r.GET("/goRegister", goRegister)
	r.POST("/register", register)
	r.GET("index", index)

	err := r.Run(":3000")
	if err != nil {
		return
	}
}
