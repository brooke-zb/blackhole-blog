package v2

import (
	"blackhole-blog/middleware/auth"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AccountLogin(c *gin.Context) {
	// bindings
	loginBody := dto.LoginDto{}
	util.BindJSON(c, &loginBody)

	// check username and password
	uid := service.User.CheckUser(loginBody.Username, loginBody.Password)
	var expire time.Duration
	if loginBody.RememberMe {
		expire = setting.Config.Server.Jwt.RememberMeExp
	} else {
		expire = setting.Config.Server.Jwt.Expire
	}

	// generate jwt token
	token, err := util.GenerateToken(uid, expire)
	if err != nil {
		log.Err.Error(err.Error())
		panic(service.NewError(service.InternalError, service.InternalErrorMessage))
	}

	util.SetCookie(c, auth.TokenKey, token, int(expire.Seconds()), true)
	c.JSON(http.StatusOK, util.RespOK(token))
}

func AccountLogout(c *gin.Context) {
	auth.MustGetUser(c)

	// remove cookie
	util.SetCookie(c, auth.TokenKey, "", -1, true)
	c.JSON(http.StatusOK, util.RespMsg("退出登录成功"))
}

func AccountInfo(c *gin.Context) {
	user := auth.MustGetUser(c)
	c.JSON(http.StatusOK, util.RespOK(user))
}

func AccountUpdateInfo(c *gin.Context) {
	user := auth.MustGetUser(c)

	// bindings
	updateInfoBody := dto.UserUpdateInfoDto{}
	util.BindJSON(c, &updateInfoBody)

	// update user info
	service.User.UpdateInfo(user.Uid, updateInfoBody)

	c.JSON(http.StatusOK, util.RespMsg("更新用户信息成功"))
}

func AccountUpdatePassword(c *gin.Context) {
	user := auth.MustGetUser(c)

	// bindings
	updatePasswordBody := dto.UserUpdatePasswordDto{}
	util.BindJSON(c, &updatePasswordBody)

	// update password
	service.User.UpdatePassword(user.Uid, updatePasswordBody)

	c.JSON(http.StatusOK, util.RespMsg("更新密码成功"))
}
