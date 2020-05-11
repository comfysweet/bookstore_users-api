package users

import (
	"github.com/comfysweet/bookstore_users-api/domain/users"
	"github.com/comfysweet/bookstore_users-api/services"
	"github.com/comfysweet/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		resultErr := errors.NewBadRequestError("invalid json body")
		c.JSON(resultErr.Status, resultErr)
		return
	}
	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user, err := services.GetUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, user)
}
