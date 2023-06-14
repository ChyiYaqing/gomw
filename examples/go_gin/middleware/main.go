package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New() // Creates a router without any middleware by default

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 作用于某个组
	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
	}
	r.Run("0.0.0.0:1207") // listen and serve on 0.0.0.0:1207
}
