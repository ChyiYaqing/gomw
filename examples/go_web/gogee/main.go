package main

import (
	"fakegee"
	"net/http"
)

func main() {
	r := fakegee.New()
	// GET() 方法添加路由
	r.GET("/", func(c *fakegee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Geek</h1>")
	})

	r.GET("/hello", func(c *fakegee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *fakegee.Context) {
		c.JSON(http.StatusOK, fakegee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":1207")
}
