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
