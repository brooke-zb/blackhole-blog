package service

import (
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/cache"
	"blackhole-blog/pkg/dao"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
	panicNotFoundErrIfNotNil(daoErr, "未找到该用户")
	return dto.ToUserDto(user)
}

func (userService) CheckUser(username string, password string) uint64 {
	user, daoErr := dao.User.FindByName(username)
	panicNotFoundErrIfNotNil(daoErr, "未找到该用户")

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

func (userService) UpdateInfo(id uint64, updateInfoBody dto.UserUpdateInfoDto) {
	// cache
	cacheKey := fmt.Sprintf("user:%d", id)
	cache.User.Delete(cacheKey)

	// update user info
	daoErr := dao.User.UpdateInfo(id, updateInfoBody)
	panicErrIfNotNil(daoErr)
}

func (userService) UpdatePassword(id uint64, updatePasswordBody dto.UserUpdatePasswordDto) {
	// cache
	cacheKey := fmt.Sprintf("user:%d", id)
	cache.User.Delete(cacheKey)

	// check password
	user, daoErr := dao.User.FindById(id)
	panicNotFoundErrIfNotNil(daoErr, "未找到该用户")
	hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(updatePasswordBody.OldPassword))
	if hashErr != nil {
		panic(NewError(Unauthorized, "密码错误"))
	}

	// hash password
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(updatePasswordBody.NewPassword), bcrypt.DefaultCost)

	// update password
	daoErr = dao.User.UpdatePassword(id, string(hashedPassword))
	panicErrIfNotNil(daoErr)
}
