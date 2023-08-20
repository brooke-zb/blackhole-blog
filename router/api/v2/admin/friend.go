package admin

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FriendFindById(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// query
	c.JSON(http.StatusOK, util.RespOK(service.Friend.FindById(param.Id)))
}

func FriendFindList(c *gin.Context) {
	// query
	friends := service.Friend.FindList()
	c.JSON(http.StatusOK, util.RespOK(friends))
}

func FriendAdd(c *gin.Context) {
	// bindings
	friend := dto.FriendAddDto{}
	util.BindJSON(c, &friend)

	// add friend
	service.Friend.Add(friend)
	c.JSON(http.StatusOK, util.RespMsg("添加成功"))
}

func FriendUpdate(c *gin.Context) {
	// bindings
	friend := dto.FriendUpdateDto{}
	util.BindJSON(c, &friend)

	// update friend
	service.Friend.Update(friend)
	c.JSON(http.StatusOK, util.RespMsg("修改成功"))
}

func FriendDelete(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// delete friend
	service.Friend.Delete(param.Id)
	c.JSON(http.StatusOK, util.RespMsg("删除成功"))
}
