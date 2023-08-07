package dto

import (
	"blackhole-blog/models"
	"time"
)

type CommentDto struct {
	Coid      uint64       `json:"coid"`
	Nickname  string       `json:"nickname"`
	Avatar    *string      `json:"avatar"`
	Site      *string      `json:"site"`
	Content   string       `json:"content"`
	CreatedAt time.Time    `json:"createdAt"`
	Children  []CommentDto `json:"children"`
	ParentId  *uint64      `json:"parentId"`
	ReplyId   *uint64      `json:"replyId"`
	ReplyTo   *string      `json:"replyTo"`
}

type CommentAddDto struct {
	Aid      uint64  `json:"aid" binding:"required"`
	Nickname string  `json:"nickname" binding:"required"`
	Content  string  `json:"content" binding:"required"`
	ReplyId  *uint64 `json:"replyId"`
	Uid      *uint64 `json:"-"`
	Ip       string  `json:"-"`
}

func ToCommentDtoList(comments models.Page[models.Comment]) (res models.Page[CommentDto]) {
	commentListDto := models.Page[CommentDto]{
		Total: comments.Total,
		Page:  comments.Page,
		Size:  comments.Size,
		Data:  make([]CommentDto, len(comments.Data)),
	}
	for i, comment := range comments.Data {
		commentListDto.Data[i] = ToCommentDto(comment)
	}
	return commentListDto
}

func ToCommentDto(comment models.Comment) CommentDto {
	commentDto := CommentDto{
		Coid:      comment.Coid,
		Nickname:  comment.Nickname,
		Avatar:    comment.Avatar,
		Site:      comment.Site,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		Children:  make([]CommentDto, len(comment.Children)),
		ParentId:  comment.ParentId,
		ReplyId:   comment.ReplyId,
		ReplyTo:   comment.ReplyTo,
	}
	for i, child := range comment.Children {
		commentDto.Children[i] = ToCommentDto(child)
	}
	return commentDto
}

func ToComment(commentAddDto CommentAddDto) models.Comment {
	return models.Comment{
		Aid:      commentAddDto.Aid,
		Nickname: commentAddDto.Nickname,
		Content:  commentAddDto.Content,
		ReplyId:  commentAddDto.ReplyId,
		Uid:      commentAddDto.Uid,
		Ip:       commentAddDto.Ip,
	}
}