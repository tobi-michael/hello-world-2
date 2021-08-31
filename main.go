package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	// create a new gin router
	router := gin.Default()

	// define a single endpoint
	router.POST("mich", helloWorldhandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	_ = router.Run(":"+ port)
}


	// run the server on the port 3000

func helloWorldhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Scooby doo",
	})
}