package dao

import "blackhole-blog/models"

type commentDao struct{}

func (commentDao) FindById(coid uint64) (comment models.Comment, err error) {
	err = db.Model(&models.Comment{}).Take(&comment, coid).Error
	return
}

func (commentDao) FindList(clause models.CommentClause) (comments models.Page[models.Comment], err error) {
	comments.Page = clause.Page
	comments.Size = clause.Size
	tx := db.Model(&models.Comment{}).Preload("Children").Where("parent_id is null")

	// 是否过滤敏感字段
	if clause.OmitSensitiveFields {
		tx = tx.Omit("email", "ip")
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
		Limit(clause.Size).Offset((clause.Page - 1) * clause.Size).
		Find(&comments.Data).Error
	return
}
