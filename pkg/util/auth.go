package util

import (
	"blackhole-blog/middleware"
	"blackhole-blog/models/dto"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
)

func ShouldGetUser(c *gin.Context) (user dto.UserDto, exists bool) {
	userObj, exists := c.Get(middleware.AuthUserKey)
	if !exists {
		return user, false
	}
	return userObj.(dto.UserDto), true
}

func MustGetUser(c *gin.Context) dto.UserDto {
	user, exists := ShouldGetUser(c)
	if !exists {
		panic(service.NewError(service.Unauthorized, service.UnauthorizedMessage))
	}
	return user
}
