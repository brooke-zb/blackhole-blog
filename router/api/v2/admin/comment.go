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

func CommentFindById(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// query
	c.JSON(http.StatusOK, util.RespOK(service.Comment.FindById(param.Id)))
}

func CommentFindList(c *gin.Context) {
	// bindings
	clause := models.CommentClause{}
	util.BindQuery(c, &clause)

	// query
	comments := service.Comment.FindList(clause)
	c.JSON(http.StatusOK, util.RespOK(comments))
}

func CommentUpdate(c *gin.Context) {
	// bindings
	comment := dto.CommentUpdateDto{}
	util.BindJSON(c, &comment)

	// update comment
	service.Comment.Update(comment)
	c.JSON(http.StatusOK, util.RespMsg("修改成功"))
}

func CommentDeleteBatch(c *gin.Context) {
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
	count := service.Comment.DeleteBatch(idList...)
	c.JSON(http.StatusOK, util.RespMsg("共删除了"+strconv.FormatInt(count, 10)+"条数据"))
}
