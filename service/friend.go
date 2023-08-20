package service

import (
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/dao"
	"blackhole-blog/pkg/util"
	"net/http"
)

type friendService struct{}

func (friendService) FindById(id uint64) dto.FriendDto {
	friend, daoErr := dao.Friend.FindById(id)
	panicNotFoundErrIfNotNil(daoErr, "未找到该友链")
	return dto.ToFriendDto(friend)
}

func (friendService) FindList() []dto.FriendDto {
	friends, daoErr := dao.Friend.FindList()
	panicErrIfNotNil(daoErr)
	return dto.ToFriendDtoList(friends)
}

func (friendService) Add(friend dto.FriendAddDto) {
	daoErr := dao.Friend.Add(friend.ToFriendModel())
	panicErrIfNotNil(daoErr)
}

func (friendService) Update(friend dto.FriendUpdateDto) {
	daoErr := dao.Friend.Update(friend)
	panicNotFoundErrIfNotNil(daoErr, "未找到该友链")
}

func (friendService) Delete(id uint64) {
	affects, daoErr := dao.Friend.Delete(id)
	panicErrIfNotNil(daoErr)
	if affects == 0 {
		panic(util.NewError(http.StatusBadRequest, "未找到该友链"))
	}
}
