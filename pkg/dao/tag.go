package dao

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"gorm.io/gorm"
)

type tagDao struct{}

func (tagDao) FindById(tid uint64) (tag models.Tag, err error) {
	err = db.Take(&tag, tid).Error
	return
}

func (tagDao) FindList(page, size int) (tags models.Page[models.Tag], err error) {
	tags.Page = page
	tags.Size = size
	err = db.Model(&models.Tag{}).
		Count(&tags.Total).
		Limit(size).Offset((page - 1) * size).
		Find(&tags.Data).Error
	return
}

func (tagDao) FindListWithHeat() (tags []models.Tag, err error) {
	err = db.Model(&models.Tag{}).Joins("inner join bh_tag_relation on bh_tag.tid = bh_tag_relation.tid").
		Select("bh_tag.tid, bh_tag.name, COUNT(*) as ArticleCount").
		Group("tid").Find(&tags).Error
	return
}

func (tagDao) Add(tag dto.TagAddDto) error {
	return db.Create(&models.Tag{Name: tag.Name}).Error
}

func (tagDao) Update(tag dto.TagUpdateDto) error {
	var count int64
	if err := db.Model(&models.Tag{}).Where("tid = ?", tag.Tid).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Model(&models.Tag{}).Where("tid = ?", tag.Tid).Update("name", tag.Name).Error
}

func (tagDao) DeleteBatch(tids ...uint64) (int64, error) {
	tx := db.Delete(&models.Tag{}, tids)
	return tx.RowsAffected, tx.Error
}
