package service

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/cache"
	"blackhole-blog/pkg/dao"
	"blackhole-blog/pkg/util"
	"net/http"
)

type roleService struct{}

func (roleService) FindById(id uint64) dto.RoleDto {
	role, daoErr := dao.Role.FindById(id)
	panicSelectErrIfNotNil(daoErr, "未找到该角色")
	return dto.ToRoleDto(role)
}

func (roleService) FindList(page, size int) models.Page[dto.RoleDto] {
	roles, daoErr := dao.Role.FindList(page, size)
	panicErrIfNotNil(daoErr)
	return dto.ToRoleDtoList(roles)
}

func (roleService) Add(role dto.RoleAddDto) {
	daoErr := dao.Role.Add(role.ToRoleModel())
	panicErrIfNotNil(daoErr, entryErr(1062, "角色名已存在"))
}

func (roleService) Update(role dto.RoleUpdateDto) {
	// cache
	defer cache.User.Clear()

	daoErr := dao.Role.Update(role)
	panicErrIfNotNil(daoErr, entryErr(1062, "角色名已存在"))
}

func (roleService) Delete(id uint64) {
	// cache
	defer cache.User.Clear()

	affects, daoErr := dao.Role.Delete(id)
	panicErrIfNotNil(daoErr)
	if affects == 0 {
		panic(util.NewError(http.StatusBadRequest, "未找到该角色"))
	}
}
