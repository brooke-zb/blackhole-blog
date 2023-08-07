package service

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/dao"
	"blackhole-blog/pkg/util"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

var (
	statusPublished = "PUBLISHED"
	statusReview    = "REVIEW"
	statusHidden    = "HIDDEN"
)

type commentService struct{}

func (commentService) FindListByArticleId(articleId uint64, page int, size int) (res models.Page[dto.CommentDto]) {
	comments, daoErr := dao.Comment.FindList(models.CommentClause{
		Aid:    &articleId,
		Status: &statusPublished,
		Page:   page,
		Size:   size,
	})
	panicErrIfNotNil(daoErr)
	return dto.ToCommentDtoList(comments)
}

func (commentService) Insert(comment models.Comment) {
	// 查询文章是否存在
	exist, daoErr := dao.Article.VerifyExist(comment.Aid)
	panicErrIfNotNil(daoErr)
	if !exist {
		panic(util.NewError(http.StatusBadRequest, "评论的文章不存在"))
	}

	// 子评论处理
	if comment.ReplyId != nil {
		reply, daoErr := dao.Comment.FindById(*comment.ReplyId)

		// 评论不存在 或者 评论非审核通过 或者 文章id不匹配
		if errors.Is(daoErr, gorm.ErrRecordNotFound) || reply.Status != statusPublished || reply.Aid != comment.Aid {
			panic(util.NewError(http.StatusBadRequest, "回复的评论不存在"))
		}

		// 评论层级判断
		if reply.ParentId != nil {
			comment.ParentId = reply.ParentId
			comment.ReplyId = &reply.Coid
			comment.ReplyTo = &reply.Nickname
		} else {
			comment.ParentId = &reply.Coid
			comment.ReplyId = nil
		}
	}

	// 检查评论者是否受信任
	comment.Status = statusPublished
	if comment.Uid == nil {
		match, _ := util.WordsFilter.FindIn(comment.Content)
		if !match {
			match, _ = util.WordsFilter.FindIn(comment.Nickname)
		}
		if match {
			comment.Status = statusReview
		}
	}

	// 生成评论id
	comment.Coid = util.NextId()

	err := dao.Comment.Insert(comment)
	panicErrIfNotNil(err)
}