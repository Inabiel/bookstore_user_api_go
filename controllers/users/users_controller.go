package users

import (
	"fmt"
	"net/http"

	"github.com/Inabiel/bookstore_user_api_go/domain/users"
	"github.com/Inabiel/bookstore_user_api_go/services"
	"github.com/Inabiel/bookstore_user_api_go/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid JSON body"))
		return
	}
	res, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
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
