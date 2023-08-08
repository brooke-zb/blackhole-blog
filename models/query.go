package models

type ArticleClause struct {
	PageParam
	Category *string `form:"category"`
	Tag      *string `form:"tag"`
	Status   *string `form:"status"`
}

type CommentClause struct {
	PageParam
	Aid                 *uint64 `form:"aid"`
	IP                  *string `form:"ip"`
	Nickname            *string `form:"nickname"`
	Status              *string `form:"status"`
	OmitSensitiveFields bool    `form:"-"`
}

type UserClause struct {
	PageParam
	Name    *string `form:"name"`
	Enabled *bool   `form:"enabled"`
}
