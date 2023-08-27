package dto

import (
	"blackhole-blog/models"
	"time"
)

type CommentDto struct {
	Coid      uint64               `json:"coid"`
	Nickname  string               `json:"nickname"`
	Uid       *uint64              `json:"uid"`
	Email     *string              `json:"email,omitempty"`
	Avatar    *string              `json:"avatar"`
	Site      *string              `json:"site"`
	Ip        string               `json:"ip,omitempty"`
	Content   string               `json:"content"`
	CreatedAt time.Time            `json:"createdAt"`
	Children  []CommentDto         `json:"children"`
	ParentId  *uint64              `json:"parentId"`
	ReplyId   *uint64              `json:"replyId"`
	ReplyTo   *string              `json:"replyTo"`
	Status    models.CommentStatus `json:"status"`
}

type CommentAddDto struct {
	Aid      uint64  `json:"aid" binding:"required"`
	Nickname string  `json:"nickname" binding:"required,min=2,max=32"`
	Content  string  `json:"content" binding:"required,max=1024"`
	ReplyId  *uint64 `json:"replyId"`
	Uid      *uint64 `json:"-"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Site     *string `json:"site" binding:"omitempty,max=200,url"`
	Ip       string  `json:"-"`
}

type CommentUpdateDto struct {
	Coid     uint64               `json:"coid" binding:"required" gorm:"-"`
	Nickname string               `json:"nickname" binding:"required,min=2,max=32"`
	Content  string               `json:"content" binding:"required,max=1024"`
	Site     *string              `json:"site" binding:"omitempty,max=200,startswith=http://|startswith=https://,url"`
	Status   models.CommentStatus `json:"status" binding:"required,oneof=PUBLISHED REVIEW HIDDEN"`
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
		Uid:       comment.Uid,
		Email:     comment.Email,
		Avatar:    comment.Avatar,
		Site:      comment.Site,
		Ip:        comment.Ip,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		Children:  make([]CommentDto, len(comment.Children)),
		ParentId:  comment.ParentId,
		ReplyId:   comment.ReplyId,
		ReplyTo:   comment.ReplyTo,
		Status:    comment.Status,
	}
	for i, child := range comment.Children {
		commentDto.Children[i] = ToCommentDto(child)
	}
	return commentDto
}

func (c CommentAddDto) ToCommentModel() models.Comment {
	return models.Comment{
		Aid:      c.Aid,
		Nickname: c.Nickname,
		Content:  c.Content,
		Site:     c.Site,
		Email:    c.Email,
		ReplyId:  c.ReplyId,
		Uid:      c.Uid,
		Ip:       c.Ip,
	}
}
