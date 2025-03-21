package models

import "time"

type Article struct {
	Aid         uint64 `gorm:"primaryKey;autoIncrement:false"`
	Uid         uint64
	User        User `gorm:"foreignKey:Uid;references:Uid"`
	Cid         uint64
	Category    Category    `gorm:"foreignKey:Cid;references:Cid"`
	Tags        []Tag       `gorm:"many2many:tagRelation;joinForeignKey:Aid;joinReferences:Tid"`
	TagRelation TagRelation `gorm:"foreignKey:Aid;references:Aid"`
	Title       string
	Abstract    *string
	Content     string
	Commentable bool
	Comments    []Comment `gorm:"foreignKey:Aid;references:Aid"`
	CreatedAt   time.Time
	UpdatedAt   *time.Time `gorm:"autoUpdateTime:false"`
	// PUBLISHED, DRAFT or HIDDEN
	Status    ArticleStatus
	ReadCount int
}

type ArticleStatus string

var (
	StatusArticlePublished ArticleStatus = "PUBLISHED"
	StatusArticleDraft     ArticleStatus = "DRAFT"
	StatusArticleHidden    ArticleStatus = "HIDDEN"
)

// TODO pin article
type PinArticle struct {
	Aid     uint64  `gorm:"primaryKey;autoIncrement:false"`
	Article Article `gorm:"foreignKey:Aid;references:Aid"`
	Weight  int
}

type Category struct {
	Cid          uint64 `gorm:"primaryKey"`
	Name         string
	Articles     []Article `gorm:"foreignKey:Cid;references:Cid"`
	ArticleCount int       `gorm:"<-:false"`
}

type Tag struct {
	Tid          uint64 `gorm:"primaryKey"`
	Name         string
	Articles     []Article `gorm:"many2many:tagRelation;joinForeignKey:Tid;joinReferences:Aid"`
	ArticleCount int       `gorm:"<-:false"`
}

type TagRelation struct {
	Aid     uint64   `gorm:"primaryKey;autoIncrement:false"`
	Tid     uint64   `gorm:"primaryKey;autoIncrement:false"`
	Article *Article `gorm:"foreignKey:Aid;references:Aid"`
	Tag     *Tag     `gorm:"foreignKey:Tid;references:Tid"`
}

type Comment struct {
	Coid      uint64 `gorm:"primaryKey;autoIncrement:false"`
	Aid       uint64
	Article   Article `gorm:"foreignKey:Aid;references:Aid"`
	Uid       *uint64
	Nickname  string
	Email     *string
	Avatar    *string `gorm:"<-:false"`
	Site      *string
	Ip        string
	Content   string
	CreatedAt time.Time
	// PUBLISHED, REVIEW or HIDDEN
	Status   CommentStatus
	Children []Comment `gorm:"foreignkey:ParentId"`
	ParentId *uint64
	ReplyId  *uint64
	ReplyTo  *string
}

type CommentStatus string

var (
	StatusCommentPublished CommentStatus = "PUBLISHED"
	StatusCommentReview    CommentStatus = "REVIEW"
	StatusCommentHidden    CommentStatus = "HIDDEN"
)

type Role struct {
	Rid         uint64 `gorm:"primaryKey"`
	Name        string
	Permissions []RolePermission `gorm:"foreignKey:Rid"`
	Users       []User           `gorm:"foreignKey:Rid"`
}

type RolePermission struct {
	Rid  uint64 `gorm:"primaryKey;autoIncrement:false"`
	Name string `gorm:"primaryKey"`
}

type User struct {
	Uid      uint64 `gorm:"primaryKey"`
	Rid      uint64
	Role     Role `gorm:"foreignKey:Rid;references:Rid"`
	Name     string
	Password string
	Mail     string
	Link     *string
	Enabled  bool
	Articles []Article `gorm:"foreignKey:Uid;references:Uid"`
}

type Friend struct {
	Fid         uint64 `gorm:"primaryKey"`
	Name        string
	Link        string
	Avatar      string
	Description *string
}
