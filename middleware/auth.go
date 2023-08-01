package middleware

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
)

const (
	AuthUserKey  = "bhblog.user"
	AuthTokenKey = "Authorization"
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
				util.SetCookie(c, AuthTokenKey, "", -1, true)
			}
		}()
		// get token from cookie
		token, err := c.Cookie(AuthTokenKey)

		// if not exist then get token from header
		if err != nil {
			token = c.GetHeader(AuthTokenKey)
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
		c.Set(AuthUserKey, user)

		// refresh token if about to expire
		if claims.ExpiresAt.Time.Sub(claims.IssuedAt.Time) < setting.Config.Server.Jwt.RefreshBeforeExp && c.Request.URL.Path != "/account/token" {
			newToken, err := util.GenerateToken(claims.Uid, setting.Config.Server.Jwt.Expire)
			if err != nil {
				log.Err.Error(err.Error())
				return
			}
			util.SetCookie(c, AuthTokenKey, newToken, int(setting.Config.Server.Jwt.Expire.Seconds()), true)
		}
	}
}
