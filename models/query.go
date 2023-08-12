package models

type ArticleClause struct {
	PageParam
	Category *string        `form:"category"`
	Tag      *string        `form:"tag"`
	Status   *ArticleStatus `form:"status" binding:"omitempty,oneof=PUBLISHED DRAFT HIDDEN"`
	Username *string        `form:"username"`
	Title    *string        `form:"title"`
	// created_at, read_count
	SortBy *string `form:"sortBy" binding:"omitempty,oneof=created_at read_count"`
}

func (c ArticleClause) Order() string {
	if c.SortBy == nil {
		return "created_at"
	}
	return *c.SortBy
}

type CommentClause struct {
	PageParam
	Aid                 *uint64        `form:"aid"`
	IP                  *string        `form:"ip" binding:"omitempty,ip"`
	Nickname            *string        `form:"nickname"`
	Status              *CommentStatus `form:"status" binding:"omitempty,oneof=PUBLISHED REVIEW HIDDEN"`
	OmitSensitiveFields bool           `form:"-" binding:"-"`
	SelectChildren      bool           `form:"-" binding:"-"`
}

type UserClause struct {
	PageParam
	Name    *string `form:"name"`
	Enabled *bool   `form:"enabled"`
}
