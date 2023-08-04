package dao

import (
	"blackhole-blog/models"
)

type articleDao struct{}

func (articleDao) FindById(uid uint64) (article models.Article, err error) {
	err = db.Model(&models.Article{}).Preload("Tags").Preload("Category").Take(&article, uid).Error
	return
}

func (articleDao) FindPreviewList(clause models.ArticleClause) (articles models.Page[models.Article], err error) {
	articles.Page = clause.Page
	articles.Size = clause.Size
	q := db.Model(&models.Article{}).Preload("Tags").Preload("Category").
		Omit("Uid", "Content", "UpdatedAt", "Status").
		Order("created_at desc")

	// 根据分类名查询
	if clause.Category != nil {
		q = q.InnerJoins("Category").
			Where("Category.name = ?", *clause.Category)
	}

	// 根据标签名查询
	if clause.Tag != nil {
		q = q.InnerJoins("TagRelation").
			Where("TagRelation.tid = (?)", db.Model(&models.Tag{}).Select("tid").Where("name = ?", *clause.Tag))
	}

	err = q.Count(&articles.Total).
		Limit(clause.Size).Offset((clause.Page - 1) * clause.Size).
		Find(&articles.Data).Error
	return
}
