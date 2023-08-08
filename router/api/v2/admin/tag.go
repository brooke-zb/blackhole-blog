package admin

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func TagFindById(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// query
	c.JSON(http.StatusOK, util.RespOK(service.Tag.FindById(param.Id)))
}

func TagFindList(c *gin.Context) {
	// bindings
	param := models.PageParam{}
	util.BindQuery(c, &param)

	// query
	c.JSON(http.StatusOK, util.RespOK(service.Tag.FindList(param.Page(), param.Size())))
}

func TagAdd(c *gin.Context) {
	// bindings
	tag := dto.TagAddDto{}
	util.BindJSON(c, &tag)

	// add tag
	service.Tag.Add(tag)
	c.JSON(http.StatusOK, util.RespMsg("添加成功"))
}

func TagUpdate(c *gin.Context) {
	// bindings
	tag := dto.TagUpdateDto{}
	util.BindJSON(c, &tag)

	// update tag
	service.Tag.Update(tag)
	c.JSON(http.StatusOK, util.RespMsg("修改成功"))
}

func TagDeleteBatch(c *gin.Context) {
	// bindings
	ids := c.Param("ids")
	if ids == "/" {
		panic(util.NewError(http.StatusUnprocessableEntity, "缺少id参数列表"))
	}

	// convert string to uint64
	rawIdList := strings.Split(ids[1:], ",")
	idList := make([]uint64, len(rawIdList))
	for i, rawId := range rawIdList {
		id, err := strconv.ParseUint(rawId, 10, 64)
		if err != nil {
			panic(util.NewError(http.StatusUnprocessableEntity, "id参数列表错误"))
		}
		idList[i] = id
	}

	// delete tag
	count := service.Tag.DeleteBatch(idList...)
	c.JSON(http.StatusOK, util.RespMsg("共删除了"+strconv.FormatInt(count, 10)+"条数据"))
}
