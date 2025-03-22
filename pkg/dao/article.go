package dao

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"

	"gorm.io/gorm"
)

type articleDao struct{}

func (articleDao) VerifyExist(aid uint64) (exist bool, err error) {
	var count int64
	err = db.Model(&models.Article{}).Where("aid = ?", aid).
		Count(&count).Error
	return count > 0, err
}

func (articleDao) FindById(aid uint64) (article models.Article, err error) {
	err = db.Model(&models.Article{}).
		Preload("Tags").Preload("Category").Preload("User").
		Take(&article, aid).Error
	return
}

func (articleDao) FindPreviewList(clause models.ArticleClause) (articles models.Page[models.Article], err error) {
	articles.Page = clause.Page()
	articles.Size = clause.Size()
	tx := db.Model(&models.Article{}).
		Preload("User").Preload("Tags").Preload("Category").
		Omit("Content", "Abstract", "UpdatedAt")

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

	// 根据用户名查询
	if clause.Username != nil {
		tx = tx.InnerJoins("User").
			Where("User.username = ?", *clause.Username)
	}

	// 根据标题模糊查询
	if clause.Title != nil {
		tx = tx.Where("title LIKE ?", "%"+*clause.Title+"%")
	}

	// 根据文章状态查询
	if clause.Status != nil {
		tx = tx.Where("status = ?", *clause.Status)
	}

	err = tx.Count(&articles.Total).
		Limit(clause.Size()).Offset((clause.Page() - 1) * clause.Size()).
		Order(clause.Order() + " desc").
		Find(&articles.Data).Error
	return
}

func (articleDao) UpdateReadCount(aid uint64, incr int64) (err error) {
	return db.Model(&models.Article{}).Where("aid = ?", aid).
		Update("read_count", gorm.Expr("read_count + ?", incr)).Error
}

func (articleDao) Add(article models.Article) (err error) {
	// 直接使用db.Create的话文章标签关系会先于文章插入，导致外键约束失败
	// 不知道是使用不规范还是gorm的bug，总之这里自行处理多对多插入
	return db.Transaction(func(tx *gorm.DB) error {
		// 插入文章
		if err := tx.Omit("Tags").Create(&article).Error; err != nil {
			return err
		}

		if len(article.Tags) > 0 {
			// 创建标签
			if err := tx.Save(&article.Tags).Error; err != nil {
				return err
			}

			// 插入标签关系
			tagNames := make([]string, len(article.Tags))
			for i, tag := range article.Tags {
				tagNames[i] = tag.Name
			}
			if err := tx.Exec("INSERT INTO bh_tag_relation(aid, tid) select ?, tid FROM bh_tag WHERE name in ?", article.Aid, tagNames).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (articleDao) Update(article dto.ArticleUpdateDto) error {
	var count int64
	if err := db.Model(&models.Article{}).Where("aid = ?", article.Aid).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return db.Transaction(func(tx *gorm.DB) error {
		// 更新文章
		if err := tx.Model(&models.Article{}).Where("aid = ?", article.Aid).Updates(article).Error; err != nil {
			return err
		}

		// 删除原有标签关系
		if err := tx.Where("aid = ?", article.Aid).Delete(&models.TagRelation{}).Error; err != nil {
			return err
		}

		if len(article.Tags) > 0 {
			// 创建标签
			if err := tx.Save(article.TagsModel()).Error; err != nil {
				return err
			}

			// 插入标签关系
			tagNames := make([]string, len(article.Tags))
			for i, tag := range article.Tags {
				tagNames[i] = tag.Name
			}
			if err := tx.Exec("INSERT INTO bh_tag_relation(aid, tid) select ?, tid FROM bh_tag WHERE name in ?", article.Aid, tagNames).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (articleDao) Delete(aid uint64) (int64, error) {
	tx := db.Where("aid = ?", aid).Delete(&models.Article{})
	return tx.RowsAffected, tx.Error
}
