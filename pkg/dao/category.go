package dao

import "blackhole-blog/models"

type categoryDao struct{}

func (categoryDao) FindListWithArticleCount() (categories []models.Category, err error) {
	// gorm丑陋的写法之一
	err = db.Model(&models.Category{}).Joins("left join bh_article on bh_category.cid = bh_article.cid").
		Select("bh_category.cid, bh_category.name, COUNT(*) as ArticleCount").
		Group("cid").Find(&categories).Error
	return
}
