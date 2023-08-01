package service

import (
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/cache"
	"blackhole-blog/pkg/dao"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct{}

func (userService) FindById(id uint64) (res dto.UserDto) {
	// cache
	cacheKey := fmt.Sprintf("user:%d", id)
	userCache := cache.User.Get(cacheKey)
	if userCache != nil && !userCache.Expired() {
		return userCache.Value()
	}
	defer cache.DeferredSetCacheWithRevocer(cache.User, cacheKey, &res)()

	user, daoErr := dao.User.FindById(id)
	if daoErr != nil {
		if errors.Is(daoErr, gorm.ErrRecordNotFound) {
			panic(NewError(NotFound, "未找到该用户"))
		} else {
			panic(NewError(InternalError, InternalErrorMessage))
		}
	}
	return dto.ToUserDto(user)
}

func (userService) CheckUser(username string, password string) uint64 {
	user, daoErr := dao.User.FindByName(username)
	if daoErr != nil {
		if errors.Is(daoErr, gorm.ErrRecordNotFound) {
			panic(NewError(NotFound, "未找到该用户"))
		} else {
			panic(NewError(InternalError, InternalErrorMessage))
		}
	}

	// check password
	hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if hashErr != nil {
		panic(NewError(Unauthorized, "密码错误"))
	}

	// check enabled
	if !user.Enabled {
		panic(NewError(Unauthorized, "该用户已被禁用"))
	}

	return user.Uid
}
