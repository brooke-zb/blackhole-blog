package admin

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CategoryFindById(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// query
	c.JSON(http.StatusOK, util.RespOK(service.Category.FindById(param.Id)))
}

func CategoryFindList(c *gin.Context) {
	// bindings
	param := models.PageParam{}
	util.BindQuery(c, &param)

	// query
	c.JSON(http.StatusOK, util.RespOK(service.Category.FindList(param.Page(), param.Size())))
}

func CategoryAdd(c *gin.Context) {
	// bindings
	category := dto.CategoryAddDto{}
	util.BindJSON(c, &category)

	// add category
	service.Category.Add(category)
	c.JSON(http.StatusOK, util.RespMsg("添加成功"))
}

func CategoryUpdate(c *gin.Context) {
	// bindings
	category := dto.CategoryUpdateDto{}
	util.BindJSON(c, &category)

	// update category
	service.Category.Update(category)
	c.JSON(http.StatusOK, util.RespMsg("修改成功"))
}

func CategoryDelete(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// delete category
	service.Category.Delete(param.Id)
	c.JSON(http.StatusOK, util.RespMsg("删除成功"))
}
