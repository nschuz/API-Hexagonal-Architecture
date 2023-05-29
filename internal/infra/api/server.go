package api

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors" // este es un alias a este paequte
)

// puede ser exportada
func RunServer() {

	server := gin.Default()

	//cors que utilice un middlwar
	server.Use(cors.Middleware(
		cors.Config{
			Origins:        "*",
			Methods:        "GET,POST,PUT,DELETE,OPTIONS,PATCH",
			RequestHeaders: "Origin, Authorization, Content-Type, Access-Control-Allow-Origin",
			MaxAge:         50 * time.Second,
		}))

	//Rgeister  app routes
	RegisterRoutes(server)

	server.Run()

}
