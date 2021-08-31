package main

import (
	"github.com/gin-gonic/gin"
)
git
func main() {
	// create a new gin router
	router := gin.Default()

	// define a single endpoint
	router.POST("mich", helloWorldhandler)

	// run the server on the port 3000
	_ = router.Run(":3000")
}

func helloWorldhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"wordings": "on the streets",
		"name":     "tobi",
		"data": gin.H{
			"something": "something",
		},
	})
}