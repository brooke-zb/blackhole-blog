package dao

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"gorm.io/gorm"
)

type roleDao struct{}

func (roleDao) FindById(id uint64) (role models.Role, err error) {
	err = db.Preload("Permissions").Take(&role, id).Error
	return
}

func (roleDao) FindList(page, size int) (roles models.Page[models.Role], err error) {
	roles.Page = page
	roles.Size = size
	err = db.Model(&models.Role{}).Preload("Permissions").
		Count(&roles.Total).
		Limit(size).Offset((page - 1) * size).
		Find(&roles.Data).Error
	return
}

func (roleDao) Add(role models.Role) error {
	return db.Create(&role).Error
}

func (roleDao) Update(role dto.RoleUpdateDto) error {
	var count int64
	if err := db.Model(&models.Role{}).Where("rid = ?", role.Rid).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if role.Name != nil {
			if err := tx.Model(&models.Role{}).Where("rid = ?", role.Rid).
				Update("name", *role.Name).Error; err != nil {
				return err
			}
		}
		if len(role.Permissions) != 0 {
			// 权限列表不为空
			if err := tx.Where("rid = ?", role.Rid).
				Not("name IN ?", role.PermissionNameList()).
				Delete(&models.RolePermission{}).Error; err != nil {
				return err
			}
			if err := tx.Save(role.PermissionList()).Error; err != nil {
				return err
			}
		} else {
			// 权限列表为空
			if err := tx.Where("rid = ?", role.Rid).
				Delete(&models.RolePermission{}).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (roleDao) Delete(rid uint64) (int64, error) {
	tx := db.Where("rid = ?", rid).Delete(&models.Role{})
	return tx.RowsAffected, tx.Error
}
