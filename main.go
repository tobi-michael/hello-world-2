package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)
type User struct {
	Name string  `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
	BloodType string `json:"blood_type"`
}

func main() {
	// create a new gin router
	router := gin.Default()

	// define a single endpoint
	router.POST("/", helloWorldhandler)



	// CRUD endpoints for data

	router.POST("/createUser", createUserHandler)
	router.GET("/getUser", getSingleUserHandler)
	router.GET("/getUsers", getAllUserHandler)
	router.PATCH("/updateUser", updateUserHandler)
	router.DELETE("/deleteUser", deleteuserhandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	_= router.Run(":" + port)
}


func createUserHandler(c *gin.Context) {

	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request data",
		})
		return

		c.JSON(200, gin.H{
			"message": "succesfully created user",
			"data": user,
		})
	}
}


func getAllUserHandler(c *gin.Context) {
			c.JSON(200, gin.H{
			"message": "hello world",
			})
}

func getSingleUserHandler(c *gin.Context)  {
	var user User
	user = User{
		Name: "Big tobss",
		Age: 5431,
		Email: "tmichael65@email.com",
	}
	c.JSON(200, gin.H{
		"message": "hello world",
		"data": user,
	})
}
func updateUserHandler(c *gin.Context) {
	fmt.Println("bjbj")
	c.JSON(200, gin.H{
		"message": "User updated!",
	})
}
func deleteuserhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user deleted!",
	})
}



	// run the server on the port 3000

func helloWorldhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Scooby doo",
	})
}

