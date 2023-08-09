package dao

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"gorm.io/gorm"
)

type commentDao struct{}

func (commentDao) FindById(coid uint64) (comment models.Comment, err error) {
	err = db.Take(&comment, coid).Error
	return
}

func (commentDao) FindList(clause models.CommentClause) (comments models.Page[models.Comment], err error) {
	comments.Page = clause.Page()
	comments.Size = clause.Size()
	tx := db.Model(&models.Comment{})

	// 是否过滤敏感字段
	if clause.OmitSensitiveFields {
		tx = tx.Omit("email", "ip")
	}

	// 是否查询子评论
	if clause.SelectChildren {
		// 子评论是否过滤敏感字段
		preloadClause := func(db *gorm.DB) *gorm.DB {
			if clause.OmitSensitiveFields {
				return db.Omit("email", "ip")
			}
			return db
		}
		tx = tx.Where("parent_id is null").Preload("Children", preloadClause)
	}

	// 根据文章id查询
	if clause.Aid != nil {
		tx = tx.Where("aid = ?", *clause.Aid)
	}

	// 根据ip查询
	if clause.IP != nil {
		tx = tx.Where("ip = ?", *clause.IP)
	}

	// 根据昵称查询
	if clause.Nickname != nil {
		tx = tx.Where("nickname = ?", *clause.Nickname)
	}

	// 根据评论状态查询
	if clause.Status != nil {
		tx = tx.Where("status = ?", *clause.Status)
	}

	err = tx.Count(&comments.Total).
		Limit(clause.Size()).Offset((clause.Page() - 1) * clause.Size()).
		Order("created_at desc").
		Find(&comments.Data).Error
	return
}

func (commentDao) Insert(comment models.Comment) error {
	return db.Create(&comment).Error
}

func (commentDao) Update(comment dto.CommentUpdateDto) error {
	return db.Model(&models.Comment{}).Where("coid = ?", comment.Coid).
		Updates(comment).Error
}

func (commentDao) DeleteBatch(coids ...uint64) (int64, error) {
	tx := db.Delete(&models.Comment{}, coids)
	return tx.RowsAffected, tx.Error
}
