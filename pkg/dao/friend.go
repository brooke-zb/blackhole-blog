package dao

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"gorm.io/gorm"
)

type friendDao struct{}

func (friendDao) FindById(fid uint64) (friend models.Friend, err error) {
	err = db.Take(&friend, fid).Error
	return
}

func (friendDao) FindList() (friends []models.Friend, err error) {
	err = db.Find(&friends).Error
	return
}

func (friendDao) Add(friend models.Friend) error {
	return db.Create(&friend).Error
}

func (friendDao) Update(friend dto.FriendUpdateDto) error {
	var count int64
	if err := db.Model(&models.Friend{}).Where("fid = ?", friend.Fid).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Model(&models.Friend{}).Where("fid = ?", friend.Fid).Updates(friend).Error
}

func (friendDao) Delete(fid uint64) (int64, error) {
	tx := db.Where("fid = ?", fid).Delete(&models.Friend{})
	return tx.RowsAffected, tx.Error
}
