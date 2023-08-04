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
	Content     string
	Commentable bool
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	Status      string
	ReadCount   int
}

type PinArticle struct {
	Aid    uint64 `gorm:"primaryKey;autoIncrement:false"`
	Weight int
}

type Category struct {
	Cid  uint64 `gorm:"primaryKey"`
	Name string
}

type Tag struct {
	Tid         uint64 `gorm:"primaryKey"`
	Name        string
	Articles    []Article   `gorm:"many2many:tag_relation;joinForeignKey:Tid;joinReferences:Aid"`
	TagRelation TagRelation `gorm:"foreignKey:Tid;references:Tid"`
}

type TagRelation struct {
	Aid uint64 `gorm:"primaryKey;autoIncrement:false"`
	Tid uint64 `gorm:"primaryKey;autoIncrement:false"`
}

type Comment struct {
	Coid      uint64 `gorm:"primaryKey;autoIncrement:false"`
	Aid       uint64
	Uid       uint64
	Nickname  string
	Email     string
	Site      string
	Ip        string
	Content   string
	CreatedAt time.Time
	Status    string
	ParentId  *uint64
	Children  []Comment `gorm:"foreignkey:ParentId"`
	ReplyId   *uint64
	ReplyTo   *string
}

type Role struct {
	Rid         uint64 `gorm:"primaryKey"`
	Name        string
	Permissions []RolePermission `gorm:"foreignKey:Rid"`
	Users       []User           `gorm:"foreignKey:Rid"`
}

type RolePermission struct {
	Rid  uint64 `gorm:"primaryKey;autoIncrement:false"`
	Name string
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
}

type FriendlyLink struct {
	Fid         uint64 `gorm:"primaryKey"`
	Name        string
	Link        string
	Avatar      string
	Description *string
}
