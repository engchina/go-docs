package main

import (
	"chapter22/model"
	"database/sql"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var adminUsers = gin.H{
	"admin1": gin.H{
		"email": "admin1@oracle.com",
		"name":  "admin1",
	},
	"admin2": gin.H{
		"email": "admin2@oracle.com",
		"name":  "admin2",
	},
	"admin3": gin.H{
		"email": "admin3@oracle.com",
		"name":  "admin3",
	},
	"admin4": gin.H{
		"email": "admin4@oracle.com",
		"name":  "admin4",
	},
	"admin5": gin.H{
		"email": "admin5@oracle.com",
		"name":  "admin5",
	},
}

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

func adminHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.FullPath()
		method := ctx.Request.Method
		user := ctx.MustGet(gin.AuthUserKey).(string)
		fmt.Println("adminHandler", path, method)
		if secret, ok := adminUsers[user]; ok {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": "No secret"})
		}
	}
}

func cookieHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("cookieHandler")
		username, err := ctx.Cookie("username")
		if err != nil {
			username = "superadmin"
			ctx.SetCookie("username", username, 60*60, "/", "localhost", false, true)
		}
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

//func testDB() {
//	dsn := "root:oracle@tcp(192.168.31.45:3306)/oracle"
//	db, err := sql.Open("mysql", dsn)
//	if err != nil {
//		fmt.Println(err)
//	}
//	err = db.Ping()
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("connect to MySQL successfully")
//}

func insert(db *sql.DB, course *model.Course) {
	db.Begin()
	defer db.Close()
	stmt, err := db.Prepare("insert into course (cname, tid) values (?,?)")
	defer stmt.Close()
	if err != nil {
		log.Println(err)
	}
	affected, err := stmt.Exec(course.Cname, course.Tid)
	log.Println(affected)
}

func main() {

	dsn := "root:oracle@tcp(192.168.31.45:3306)/oracle"
	err := initDB(dsn)
	if err != nil {
		fmt.Println(err)
	}
	course := &model.Course{
		Cname: "user1",
		Tid:   10,
	}
	insert(dbs, course)

	r := gin.Default()
	routerGroup := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin1": "123456",
		"admin2": "123456",
		"admin3": "123456",
		"admin4": "123456",
		"admin5": "123456",
	}))
	routerGroup.GET("/secret", adminHandler())

	//r.Use(myHandler, myBaseHandler())
	r.Use(cookieHandler())
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
	r.GET("/index", index)

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			err := session.Save()
			if err != nil {
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"hello": session.Get("hello")})
	})

	err = r.Run(":3000")
	if err != nil {
		return
	}
}

var dbs *sql.DB

func initDB(dsn string) (err error) {
	dbs, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	//defer func(dbs *sql.DB) {
	//	err := dbs.Close()
	//	if err != nil {
	//		return
	//	}
	//}(dbs)

	err = dbs.Ping()
	if err != nil {
		return err
	}

	dbs.SetMaxOpenConns(10)
	dbs.SetMaxIdleConns(5)

	return nil
}
