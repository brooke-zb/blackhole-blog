package v2

import (
	"blackhole-blog/middleware/auth"
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommentFindListByArticleId(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)
	page := models.PageParam{}
	util.BindQuery(c, &page)

	// query
	comments := service.Comment.FindList(models.CommentClause{
		Aid:                 &param.Id,
		PageParam:           page,
		Status:              &models.StatusCommentPublished,
		OmitSensitiveFields: true,
		SelectChildren:      true,
	})
	c.JSON(http.StatusOK, util.RespOK(comments))
}

func CommentAdd(c *gin.Context) {
	// bindings
	comment := dto.CommentAddDto{}
	util.BindJSON(c, &comment)
	comment.Ip = c.ClientIP()
	user, exist := auth.ShouldGetUser(c)
	if exist {
		comment.Uid = &user.Uid
	}

	// insert
	status := service.Comment.Insert(comment)
	if status == models.StatusCommentReview {
		c.JSON(http.StatusOK, util.RespMsg("评论成功，等待博主审核"))
		return
	}
	c.JSON(http.StatusOK, util.RespMsg("评论成功"))
}
