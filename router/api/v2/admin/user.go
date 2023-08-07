package admin

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserFindById(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// query
	user := service.User.FindById(param.Id)
	c.JSON(http.StatusOK, util.RespOK(user))
}

func UserFindList(c *gin.Context) {
	// bindings
	clause := models.UserClause{}
	util.BindQuery(c, &clause)

	// query
	users := service.User.FindList(clause)
	c.JSON(http.StatusOK, util.RespOK(users))
}

func UserAdd(c *gin.Context) {
	// bindings
	user := dto.UserAddDto{}
	util.BindJSON(c, &user)

	// insert user
	service.User.Add(user)
	c.JSON(http.StatusOK, util.RespMsg("添加成功"))
}

func UserUpdate(c *gin.Context) {
	// bindings
	user := dto.UserUpdateDto{}
	util.BindJSON(c, &user)

	// update user
	service.User.Update(user)
	c.JSON(http.StatusOK, util.RespMsg("修改成功"))
}

func UserDelete(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// delete user
	service.User.Delete(param.Id)
	c.JSON(http.StatusOK, util.RespMsg("删除成功"))
}
