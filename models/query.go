package models

type ArticleClause struct {
	Category *string
	Tag      *string
	Page     int
	Size     int
}

type CommentClause struct {
	Aid                 *uint64
	IP                  *string
	Nickname            *string
	Status              *string
	OmitSensitiveFields bool
	Page                int
	Size                int
}
