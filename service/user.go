package service

import (
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/cache"
	"blackhole-blog/pkg/dao"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type userService struct{}

func (userService) FindById(id uint64) (res dto.UserDto) {
	// cache
	cacheKey := fmt.Sprintf("user:%d", id)
	userCache := cache.User.Get(cacheKey)
	if userCache != nil && !userCache.Expired() {
		return userCache.Value()
	}
	defer cache.DeferredSetWithRevocer(cache.User, cacheKey, &res)()

	user, daoErr := dao.User.FindById(id)
	panicSelectErrIfNotNil(daoErr, "未找到该用户")
	return dto.ToUserDto(user)
}

func (userService) CheckUser(username string, password string) uint64 {
	user, daoErr := dao.User.FindByName(username)
	panicSelectErrIfNotNil(daoErr, "未找到该用户")

	// check password
	hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if hashErr != nil {
		panic(util.NewError(http.StatusUnauthorized, "密码错误"))
	}

	// check enabled
	if !user.Enabled {
		panic(util.NewError(http.StatusUnauthorized, "该用户已被禁用"))
	}

	return user.Uid
}

func (userService) UpdateInfo(id uint64, updateInfoBody dto.UserUpdateInfoDto) {
	// cache
	cacheKey := fmt.Sprintf("user:%d", id)
	defer cache.User.Delete(cacheKey)

	// update user info
	daoErr := dao.User.UpdateInfo(id, updateInfoBody)
	panicErrIfNotNil(daoErr, entryErrProducer(1062, updateInfoErrProducer))
}

func updateInfoErrProducer(msg string) string {
	if strings.Contains(msg, "mail") {
		return "邮箱已被使用"
	} else if strings.Contains(msg, "name") {
		return "昵称已被使用"
	} else {
		return setting.InternalErrorMessage
	}
}

func (userService) UpdatePassword(id uint64, updatePasswordBody dto.UserUpdatePasswordDto) {
	// cache
	cacheKey := fmt.Sprintf("user:%d", id)
	defer cache.User.Delete(cacheKey)

	// check password
	user, daoErr := dao.User.FindById(id)
	panicSelectErrIfNotNil(daoErr, "未找到该用户")
	hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(updatePasswordBody.OldPassword))
	if hashErr != nil {
		panic(util.NewError(http.StatusUnauthorized, "密码错误"))
	}

	// hash password
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(updatePasswordBody.NewPassword), bcrypt.DefaultCost)

	// update password
	daoErr = dao.User.UpdatePassword(id, string(hashedPassword))
	panicErrIfNotNil(daoErr)
}
