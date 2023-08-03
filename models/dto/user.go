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

func ToUserDto(user models.User) UserDto {
	userDto := UserDto{
		Uid: user.Uid,
		Role: RoleDto{
			Rid:         user.Role.Rid,
			Name:        user.Role.Name,
			Permissions: make([]RolePermissionDto, len(user.Role.Permissions)),
		},
		Name:    user.Name,
		Mail:    user.Mail,
		Link:    user.Link,
		Enabled: user.Enabled,
	}
	for i, permission := range user.Role.Permissions {
		userDto.Role.Permissions[i] = RolePermissionDto{
			Name: permission.Name,
		}
	}
	return userDto
}
