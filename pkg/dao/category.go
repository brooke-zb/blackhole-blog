package dao

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
)

type categoryDao struct{}

func (categoryDao) FindById(cid uint64) (category models.Category, err error) {
	err = db.Take(&category, cid).Error
	return
}

func (categoryDao) FindList(page, size int) (categories models.Page[models.Category], err error) {
	categories.Page = page
	categories.Size = size
	err = db.Model(&models.Category{}).
		Count(&categories.Total).
		Limit(size).Offset((page - 1) * size).
		Find(&categories.Data).Error
	return
}

func (categoryDao) FindListWithArticleCount() (categories []models.Category, err error) {
	// gorm丑陋的写法之一
	err = db.Joins("LEFT JOIN bh_article ON bh_category.cid = bh_article.cid").
		Select("bh_category.cid, bh_category.name, COUNT(*) AS ArticleCount").
		Group("cid").Find(&categories).Error
	return
}

func (categoryDao) Add(category dto.CategoryAddDto) error {
	return db.Create(&models.Category{Name: category.Name}).Error
}

func (categoryDao) Update(category dto.CategoryUpdateDto) error {
	return db.Model(&models.Category{}).Where("cid = ?", category.Cid).
		Update("name", category.Name).Error
}

func (categoryDao) Delete(cid uint64) error {
	return db.Delete(&models.Category{}, cid).Error
}
