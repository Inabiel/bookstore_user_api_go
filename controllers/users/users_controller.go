package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Inabiel/bookstore_user_api_go/domain/users"
	"github.com/Inabiel/bookstore_user_api_go/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: handle any err
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//TODO: handle json decode err
		return
	}
	res, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation error
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg": res,
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"msg": "not implemented yet",
	})
}

func FindUser(c *gin.Context) {
	user := c.Param("id")
	c.JSON(http.StatusNotImplemented, gin.H{
		"msg": fmt.Sprintf("user id is %s", user),
	})
}

func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"msg": "not implemented yet",
	})
}
