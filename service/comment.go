package service

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/dao"
)

var statusPublished = "PUBLISHED"

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
