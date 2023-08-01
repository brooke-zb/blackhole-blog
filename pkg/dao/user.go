package dao

import "blackhole-blog/models"

type userDao struct{}

func (userDao) FindById(uid uint64) (user models.User, err error) {
	err = db.Preload("Role").Preload("Role.Permissions").Take(&user, uid).Error
	return
}

func (userDao) FindByName(username string) (user models.User, err error) {
	err = db.Preload("Role").Preload("Role.Permissions").Where("name = ?", username).Take(&user).Error
	return
}

func (userDao) FindList(page int, pageSize int) (users models.Page[models.User], err error) {
	users.Page = page
	users.PageSize = pageSize
	err = db.Model(&models.User{}).Preload("Role").Preload("Role.Permissions").
		Count(&users.Total).
		Limit(pageSize).Offset((page - 1) * pageSize).
		Find(&users.Data).Error
	return
}
