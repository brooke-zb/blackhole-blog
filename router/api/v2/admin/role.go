package admin

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoleFindById(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// query
	role := service.Role.FindById(param.Id)
	c.JSON(http.StatusOK, util.RespOK(role))
}

func RoleFindList(c *gin.Context) {
	// bindings
	param := models.PageParam{}
	util.BindQuery(c, &param)

	// query
	roles := service.Role.FindList(param.Page(), param.Size())
	c.JSON(http.StatusOK, util.RespOK(roles))
}

func RoleAdd(c *gin.Context) {
	// bindings
	role := dto.RoleAddDto{}
	util.BindJSON(c, &role)

	// insert role
	service.Role.Add(role)
	c.JSON(http.StatusOK, util.RespMsg("添加成功"))
}

func RoleUpdate(c *gin.Context) {
	// bindings
	role := dto.RoleUpdateDto{}
	util.BindJSON(c, &role)

	// update role
	service.Role.Update(role)
	c.JSON(http.StatusOK, util.RespMsg("修改成功"))
}

func RoleDelete(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// delete role
	service.Role.Delete(param.Id)
	c.JSON(http.StatusOK, util.RespMsg("删除成功"))
}
