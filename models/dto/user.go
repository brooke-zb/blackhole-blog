package dto

import "blackhole-blog/models"

type UserDto struct {
	Uid     uint64  `json:"uid"`
	Role    RoleDto `json:"role"`
	Name    string  `json:"name"`
	Mail    string  `json:"mail"`
	Link    *string `json:"link"`
	Enabled bool    `json:"enabled"`
}

type LoginDto struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe"`
}

type UserUpdateInfoDto struct {
	Name *string `json:"name" binding:"omitempty,min=2,max=32"`
	Mail *string `json:"mail" binding:"omitempty,email"`
	Link *string `json:"link" binding:"omitempty,url"`
}

type UserUpdatePasswordDto struct {
	OldPassword string `json:"oldPassword" binding:"required,min=6,max=32"`
	NewPassword string `json:"newPassword" binding:"required,min=6,max=32"`
}

type UserAddDto struct {
	Name     string  `json:"name" binding:"required,min=2,max=32"`
	Password string  `json:"password" binding:"required,min=6,max=32"`
	Mail     string  `json:"mail" binding:"required,email"`
	Link     *string `json:"link" binding:"omitempty,url"`
	Rid      uint64  `json:"rid" binding:"required"`
	Enabled  bool    `json:"enabled"`
}

type UserUpdateDto struct {
	Uid      uint64  `json:"uid" binding:"required" gorm:"-"`
	Name     *string `json:"name" binding:"omitempty,min=2,max=32"`
	Password *string `json:"password" binding:"omitempty,min=6,max=32"`
	Mail     *string `json:"mail" binding:"omitempty,email"`
	Link     *string `json:"link" binding:"omitempty,url"`
	Rid      *uint64 `json:"rid"`
	Enabled  *bool   `json:"enabled"`
}

func ToUserDtoList(users models.Page[models.User]) models.Page[UserDto] {
	userDtoList := models.Page[UserDto]{
		Total: users.Total,
		Page:  users.Page,
		Size:  users.Size,
		Data:  make([]UserDto, len(users.Data)),
	}
	for i, user := range users.Data {
		userDtoList.Data[i] = ToUserDto(user)
	}
	return userDtoList
}

func ToUserDto(user models.User) UserDto {
	userDto := UserDto{
		Uid:     user.Uid,
		Role:    ToRoleDto(user.Role),
		Name:    user.Name,
		Mail:    user.Mail,
		Link:    user.Link,
		Enabled: user.Enabled,
	}
	return userDto
}

func (u UserAddDto) ToUserModel() models.User {
	return models.User{
		Name:     u.Name,
		Password: u.Password,
		Mail:     u.Mail,
		Link:     u.Link,
		Rid:      u.Rid,
		Enabled:  u.Enabled,
	}
}
