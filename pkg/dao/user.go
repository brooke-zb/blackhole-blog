package dao

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"gorm.io/gorm"
)

type userDao struct{}

func (userDao) FindById(uid uint64) (user models.User, err error) {
	err = db.Preload("Role").Preload("Role.Permissions").Take(&user, uid).Error
	return
}

func (userDao) FindByName(username string) (user models.User, err error) {
	err = db.Preload("Role").Preload("Role.Permissions").Where("name = ?", username).Take(&user).Error
	return
}

func (userDao) FindList(clause models.UserClause) (users models.Page[models.User], err error) {
	users.Page = clause.Page()
	users.Size = clause.Size()
	tx := db.Model(&models.User{}).Preload("Role").Preload("Role.Permissions")

	// 根据用户名模糊查询
	if clause.Name != nil {
		tx = tx.Where("name LIKE ?", "%"+*clause.Name+"%")
	}

	// 根据状态查询
	if clause.Enabled != nil {
		tx = tx.Where("enabled = ?", *clause.Enabled)
	}

	err = tx.Count(&users.Total).
		Limit(clause.Size()).Offset((clause.Page() - 1) * clause.Size()).
		Find(&users.Data).Error
	return
}

func (userDao) UpdateInfo(uid uint64, updateInfoBody dto.UserUpdateInfoDto) error {
	return db.Model(&models.User{}).Where("uid = ?", uid).
		Updates(updateInfoBody).Error
}

func (userDao) UpdatePassword(uid uint64, hashedPassword string) error {
	return db.Model(&models.User{}).Where("uid = ?", uid).
		Update("password", hashedPassword).Error
}

func (userDao) Add(user models.User) error {
	return db.Create(&user).Error
}

func (userDao) Update(uid uint64, user dto.UserUpdateDto) error {
	var count int64
	if err := db.Model(&models.User{}).Where("uid = ?", user.Uid).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Model(&models.User{}).Where("uid = ?", uid).
		Updates(user).Error
}

func (userDao) Delete(uid uint64) (int64, error) {
	tx := db.Where("uid = ?", uid).Delete(&models.User{})
	return tx.RowsAffected, tx.Error
}
