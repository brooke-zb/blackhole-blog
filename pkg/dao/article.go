package dao

import (
	"blackhole-blog/models"
	"gorm.io/gorm"
)

type articleDao struct{}

func (articleDao) VerifyExist(aid uint64) (exist bool, err error) {
	var count int64
	err = db.Model(&models.Article{}).Where("aid = ?", aid).Count(&count).Error
	return count > 0, err
}

func (articleDao) FindById(aid uint64) (article models.Article, err error) {
	err = db.Model(&models.Article{}).Preload("Tags").Preload("Category").Take(&article, aid).Error
	return
}

func (articleDao) FindPreviewList(clause models.ArticleClause) (articles models.Page[models.Article], err error) {
	articles.Page = clause.Page()
	articles.Size = clause.Size()
	tx := db.Model(&models.Article{}).Preload("Tags").Preload("Category").
		Omit("Uid", "Content", "UpdatedAt", "Status").
		Order("created_at desc")

	// 根据分类名查询
	if clause.Category != nil {
		tx = tx.InnerJoins("Category").
			Where("Category.name = ?", *clause.Category)
	}

	// 根据标签名查询
	if clause.Tag != nil {
		tx = tx.InnerJoins("TagRelation").InnerJoins("TagRelation.Tag").
			Where("TagRelation__Tag.name = ?", *clause.Tag)
	}

	err = tx.Count(&articles.Total).
		Limit(clause.Size()).Offset((clause.Page() - 1) * clause.Size()).
		Find(&articles.Data).Error
	return
}

func (articleDao) UpdateReadCount(aid uint64, incr int64) (err error) {
	err = db.Model(&models.Article{}).Where("aid = ?", aid).Update("read_count", gorm.Expr("read_count + ?", incr)).Error
	return
}
