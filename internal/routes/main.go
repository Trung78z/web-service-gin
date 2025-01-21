package routes

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRoutes(development bool) *gin.Engine {
	if !development {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		f, _ := os.Create("pkg/logger/gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	} else {
		fmt.Println("Development mode")
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")
	addUserRoutes(v1)
	return r
}
