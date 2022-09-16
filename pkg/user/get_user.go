package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	//get user by login
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, user)

}
