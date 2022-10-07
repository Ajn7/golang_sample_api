//1.go mod init example/sample_userpass_api(root file)
//2. go get github.com/gin-gonic/gin (package)
//3.Code

package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Email    string `json:"email"`
	Password string `josn:"password"`
}

//some data
var users = []user{
	{Email: "google.com", Password: "12345"},
	{Email: "fb.com", Password: "abc@fb"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return //returns error by c.BindJSON
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
func getUserPass(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return //returns error by c.BindJSON
	}
	for _, val := range users {
		if val.Email == newUser.Email {

			if val.Password == newUser.Password {

				c.IndentedJSON(http.StatusOK, newUser)

			}
		}
	}

	c.IndentedJSON(http.StatusBadRequest, "Incorrect Pssword or Email ")

}
func userByEmail(c *gin.Context) {
	email := c.Param("email")
	password := c.Param("password")
	user, err := getUserByEmail(email, password)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User Not Found"})
	}
	c.IndentedJSON(http.StatusOK, user)
}
func getUserByEmail(email string, password string) (*user, error) {

	for i, val := range users {
		if val.Email == email {

			if val.Password == password {

				return &users[i], nil

			}
		}
	}
	return nil, errors.New("book not found")
}
func main() {

	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:email/:password", userByEmail) // format-users/email/password
	router.POST("/adduser", createUser)
	router.POST("/login", getUserPass)
	router.Run("localhost:2000")
}
