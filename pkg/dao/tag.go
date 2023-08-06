package dao

import "blackhole-blog/models"

type tagDao struct{}

func (tagDao) FindListWithHeat() (tags []models.Tag, err error) {
	err = db.Model(&models.Tag{}).Joins("left join bh_tag_relation on bh_tag.tid = bh_tag_relation.tid").
		Select("bh_tag.tid, bh_tag.name, COUNT(*) as ArticleCount").
		Group("tid").Find(&tags).Error
	return
}
