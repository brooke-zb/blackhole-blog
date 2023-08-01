package auth

import (
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
)

const (
	UserKey  = "bhblog.user"
	TokenKey = "Authorization"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			r := recover()
			if r != nil {
				err, ok := r.(service.Error)
				if ok {
					log.Err.Error(err.Error())
				}
				util.SetCookie(c, TokenKey, "", -1, true)
			}
		}()
		// get token from cookie
		token, err := c.Cookie(TokenKey)

		// if not exist then get token from header
		if err != nil {
			token = c.GetHeader(TokenKey)
			if token == "" {
				return
			}
		}

		// verify token
		claims, err := util.VerifyToken(token)
		if err != nil {
			log.Err.Error(err.Error())
			return
		}

		// get user by uid
		user := service.User.FindById(claims.Uid)

		// set user to context
		c.Set(UserKey, user)

		// refresh token if about to expire
		if claims.ExpiresAt.Time.Sub(claims.IssuedAt.Time) < setting.Config.Server.Jwt.RefreshBeforeExp && c.Request.URL.Path != "/account/token" {
			newToken, err := util.GenerateToken(claims.Uid, setting.Config.Server.Jwt.Expire)
			if err != nil {
				log.Err.Error(err.Error())
				return
			}
			util.SetCookie(c, TokenKey, newToken, int(setting.Config.Server.Jwt.Expire.Seconds()), true)
		}
	}
}

func ShouldGetUser(c *gin.Context) (user dto.UserDto, exists bool) {
	userObj, exists := c.Get(UserKey)
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
