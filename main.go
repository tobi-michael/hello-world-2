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

var Users []User

func main() {
	// create a new gin router
	router := gin.Default()

	// define a single endpoint
	router.POST("/", helloWorldhandler)



	// CRUD endpoints for data

	router.POST("/createUser", createUserHandler)
	router.GET("/getUser/:name", getSingleUserHandler)
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

	}

	Users = append(Users, user)

	c.JSON(200, gin.H{
		"message": "successfully created user",
		"data": user,
	})
}


func getAllUserHandler(c *gin.Context) {
			c.JSON(200, gin.H{
			"message": "hello world",
			"data" : Users,
			})
}

func getSingleUserHandler(c *gin.Context)  {
	name := c.Param("name")

	fmt.Println("name", name)
	var user User
	for _, value := range Users {

		// check the current iteration of users
		// check if the name matches the client request
		if value.Name == name {

			// if it matches assign the value to the empty user object we created
			user = value
		}
	}

	// if no match was found
	// the empty use we created would still be empty
	// check if user is empty, if so return a not found error

	if &user == nil {
		c.JSON(404, gin.H{
			"error": "no user with name found: " + name,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
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

