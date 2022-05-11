package users

import (
	"github.com/comfysweet/bookstore-oauth-go/oauth"
	"github.com/comfysweet/bookstore_users-api/domain/users"
	"github.com/comfysweet/bookstore_users-api/services"
	"github.com/comfysweet/bookstore_utils-go/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}
	return userID, nil
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		resultErr := errors.NewBadRequestError("invalid json body")
		c.JSON(resultErr.Status, resultErr)
		return
	}
	result, err := services.UserService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result.Marshal(oauth.IsPublic(c.Request)))
}

func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	if callerId := oauth.GetCallerId(c.Request); callerId == 0 {
		err := errors.RestErr{
			Message: "resource nor available",
			Status:  http.StatusUnauthorized,
		}
		c.JSON(err.Status, err)
		return
	}

	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	user, err := services.UserService.GetUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshal(false))
		return
	}
	c.JSON(http.StatusOK, user.Marshal(oauth.IsPublic(c.Request)))
}

func Update(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		resultErr := errors.NewBadRequestError("invalid json body")
		c.JSON(resultErr.Status, resultErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(oauth.IsPublic(c.Request)))
}

func Delete(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	result, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(oauth.IsPublic(c.Request)))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, user.Marshal(oauth.IsPublic(c.Request)))
}
