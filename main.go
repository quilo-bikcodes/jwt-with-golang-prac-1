package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quilo-bikcodes/Go-JWT/controllers"
	"github.com/quilo-bikcodes/Go-JWT/initializers"
	"github.com/quilo-bikcodes/Go-JWT/middleware"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncData()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate",middleware.RequireAuth, controllers.Validate)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}
