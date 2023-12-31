package service

import (
	"blackhole-blog/models"
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

func (userService) FindList(clause models.UserClause) models.Page[dto.UserDto] {
	users, daoErr := dao.User.FindList(clause)
	panicErrIfNotNil(daoErr)
	return dto.ToUserDtoList(users)
}

func (userService) FindById(id uint64) (res dto.UserDto) {
	// cache
	cacheKey := fmt.Sprintf("user:%d", id)
	userCache := cache.User.Get(cacheKey)
	if userCache != nil && !userCache.Expired() {
		return userCache.Value()
	}
	defer cache.DeferredSetWithRecover(cache.User, cacheKey, &res)()

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
	panicErrIfNotNil(daoErr, entryErrProducer(1062, userDuplicateErrProducer))
}

func userDuplicateErrProducer(msg string) string {
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
	panicNotFoundErrIfNotNil(daoErr, "未找到该用户")
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

func (userService) Add(user dto.UserAddDto) {
	// hash password
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	panicErrIfNotNil(hashErr)
	user.Password = string(hashedPassword)

	// insert user
	daoErr := dao.User.Add(user.ToUserModel())
	panicErrIfNotNil(daoErr, entryErrProducer(1062, userDuplicateErrProducer), entryErr(1452, "未找到该角色"))
}

func (userService) Update(user dto.UserUpdateDto) {
	// cache
	cacheKey := fmt.Sprintf("user:%d", user.Uid)
	defer cache.User.Delete(cacheKey)

	// hash password
	if user.Password != nil {
		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
		panicErrIfNotNil(hashErr)
		newPassword := string(hashedPassword)
		user.Password = &newPassword
	}

	// update user
	daoErr := dao.User.Update(user.Uid, user)
	panicNotFoundErrIfNotNil(daoErr, "未找到该用户", entryErrProducer(1062, userDuplicateErrProducer), entryErr(1452, "未找到该角色"))
}

func (userService) Delete(id uint64) {
	// cache
	cacheKey := fmt.Sprintf("user:%d", id)
	defer cache.User.Delete(cacheKey)

	// delete user
	affects, daoErr := dao.User.Delete(id)
	panicErrIfNotNil(daoErr, entryErr(1451, "该用户下存在文章，禁止删除"))
	if affects == 0 {
		panic(util.NewError(http.StatusBadRequest, "未找到该用户"))
	}
}
