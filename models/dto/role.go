package dto

import (
	"blackhole-blog/models"
)

type RoleDto struct {
	Rid         uint64              `json:"rid" binding:"required"`
	Name        string              `json:"name"`
	Permissions []RolePermissionDto `json:"permissions" binding:"required"`
}

type RolePermissionDto struct {
	Name string `json:"name"`
}

type RoleUpdateDto struct {
	Rid         uint64              `json:"rid" binding:"required" gorm:"-"`
	Name        *string             `json:"name" binding:"omitempty,min=1,max=20"`
	Permissions []RolePermissionDto `json:"permissions"`
}

func (r RoleUpdateDto) ToRoleModel() models.Role {
	role := models.Role{
		Rid:         r.Rid,
		Permissions: r.PermissionList(),
	}
	if r.Name != nil {
		role.Name = *r.Name
	}
	return role
}

func (r RoleUpdateDto) PermissionNameList() []string {
	permissions := make([]string, len(r.Permissions))
	for i, permission := range r.Permissions {
		permissions[i] = permission.Name
	}
	return permissions
}

func (r RoleUpdateDto) PermissionList() []models.RolePermission {
	permissions := make([]models.RolePermission, len(r.Permissions))
	for i, permission := range r.Permissions {
		permissions[i] = models.RolePermission{
			Rid:  r.Rid,
			Name: permission.Name,
		}
	}
	return permissions
}

type RoleAddDto struct {
	Name        string              `json:"name" binding:"required,min=1,max=20"`
	Permissions []RolePermissionDto `json:"permissions" binding:"required"`
}

func (r RoleAddDto) ToRoleModel() models.Role {
	role := models.Role{
		Name:        r.Name,
		Permissions: make([]models.RolePermission, len(r.Permissions)),
	}
	for i, permission := range r.Permissions {
		role.Permissions[i] = models.RolePermission{
			Name: permission.Name,
		}
	}
	return role
}

func ToRoleDtoList(roles models.Page[models.Role]) models.Page[RoleDto] {
	roleDtoList := models.Page[RoleDto]{
		Total: roles.Total,
		Page:  roles.Page,
		Size:  roles.Size,
		Data:  make([]RoleDto, len(roles.Data)),
	}
	for i, role := range roles.Data {
		roleDtoList.Data[i] = ToRoleDto(role)
	}
	return roleDtoList
}

func ToRoleDto(role models.Role) RoleDto {
	roleDto := RoleDto{
		Rid:         role.Rid,
		Name:        role.Name,
		Permissions: make([]RolePermissionDto, len(role.Permissions)),
	}
	for i, permission := range role.Permissions {
		roleDto.Permissions[i] = RolePermissionDto{
			Name: permission.Name,
		}
	}
	return roleDto
}
