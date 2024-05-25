package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quilo-bikcodes/Go-JWT/controllers"
	"github.com/quilo-bikcodes/Go-JWT/initializers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncData()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}
