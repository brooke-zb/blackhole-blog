package dto

import "blackhole-blog/models"

type FriendDto struct {
	Fid         uint64  `json:"fid"`
	Name        string  `json:"name"`
	Link        string  `json:"link"`
	Avatar      string  `json:"avatar"`
	Description *string `json:"description"`
}

type FriendAddDto struct {
	Name        string  `json:"name" binding:"required"`
	Link        string  `json:"link" binding:"required,url"`
	Avatar      string  `json:"avatar" binding:"required,url"`
	Description *string `json:"description"`
}

func (f FriendAddDto) ToFriendModel() models.Friend {
	return models.Friend{
		Name:        f.Name,
		Link:        f.Link,
		Avatar:      f.Avatar,
		Description: f.Description,
	}
}

type FriendUpdateDto struct {
	Fid         uint64  `json:"fid" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Link        string  `json:"link" binding:"required"`
	Avatar      string  `json:"avatar" binding:"required"`
	Description *string `json:"description"`
}

func ToFriendDto(friend models.Friend) FriendDto {
	return FriendDto{
		Fid:         friend.Fid,
		Name:        friend.Name,
		Link:        friend.Link,
		Avatar:      friend.Avatar,
		Description: friend.Description,
	}
}

func ToFriendDtoList(friends []models.Friend) []FriendDto {
	friendDtoList := make([]FriendDto, len(friends))
	for i, friend := range friends {
		friendDtoList[i] = ToFriendDto(friend)
	}
	return friendDtoList
}
