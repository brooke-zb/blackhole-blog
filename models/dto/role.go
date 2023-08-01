package dto

type RoleDto struct {
	Rid         uint64              `json:"rid"`
	Name        string              `json:"name"`
	Permissions []RolePermissionDto `json:"permissions"`
}

type RolePermissionDto struct {
	Name string `json:"name"`
}
