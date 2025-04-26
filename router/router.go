package router

import "github.com/gin-gonic/gin"

func Initialize() {

	//Initilize router
	r := gin.Default()
	//Initialize Routes
	initializeRoutes(r)
	r.Run() //run server
}
