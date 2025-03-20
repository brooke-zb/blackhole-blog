package dto

import (
	"blackhole-blog/models"
	"time"
)

type ArticleDto struct {
	Aid         uint64               `json:"aid"`
	User        ArticleUserDto       `json:"user"`
	Category    ArticleCategoryDto   `json:"category"`
	Tags        []ArticleTagDto      `json:"tags"`
	Title       string               `json:"title"`
	Abstract    *string              `json:"abstract"`
	Content     string               `json:"content"`
	Commentable bool                 `json:"commentable"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   *time.Time           `json:"updatedAt"`
	Status      models.ArticleStatus `json:"status"`
	ReadCount   int                  `json:"readCount"`
}

type ArticlePreviewDto struct {
	Aid       uint64               `json:"aid"`
	User      ArticleUserDto       `json:"user"`
	Category  ArticleCategoryDto   `json:"category"`
	Tags      []ArticleTagDto      `json:"tags"`
	Title     string               `json:"title"`
	Abstract  *string              `json:"abstract"`
	CreatedAt time.Time            `json:"createdAt"`
	Status    models.ArticleStatus `json:"status"`
	ReadCount int                  `json:"readCount"`
}

type ArticleUserDto struct {
	Uid  uint64 `json:"uid"`
	Name string `json:"name"`
}

type ArticleCategoryDto struct {
	Cid  uint64 `json:"cid"`
	Name string `json:"name"`
}

type ArticleTagDto struct {
	Name string `json:"name" binding:"required"`
}

type ArticleAddDto struct {
	Uid         uint64               `json:"-"`
	Cid         uint64               `json:"cid" binding:"required"`
	Tags        []ArticleTagDto      `json:"tags" binding:"required"`
	Title       string               `json:"title" binding:"required,max=64"`
	Abstract    *string              `json:"abstract" binding:"omitempty,max=255"`
	Content     string               `json:"content" binding:"required"`
	Commentable bool                 `json:"commentable" binding:"required"`
	Status      models.ArticleStatus `json:"status" binding:"required,oneof=PUBLISHED DRAFT HIDDEN"`
}

func (a ArticleAddDto) ToArticleModel() models.Article {
	article := models.Article{
		Uid:         a.Uid,
		Cid:         a.Cid,
		Title:       a.Title,
		Abstract:    a.Abstract,
		Tags:        make([]models.Tag, len(a.Tags)),
		Content:     a.Content,
		Commentable: a.Commentable,
		Status:      a.Status,
	}
	for i, tag := range a.Tags {
		article.Tags[i] = models.Tag{
			Name: tag.Name,
		}
	}
	return article
}

type ArticleUpdateDto struct {
	Aid         uint64               `json:"aid" binding:"required" gorm:"-"`
	Cid         uint64               `json:"cid" binding:"required"`
	Tags        []ArticleTagDto      `json:"tags" binding:"required" gorm:"-"`
	Title       string               `json:"title" binding:"required,max=64"`
	Abstract    *string              `json:"abstract" binding:"omitempty,max=255"`
	Content     string               `json:"content" binding:"required"`
	Commentable bool                 `json:"commentable" binding:"required"`
	Status      models.ArticleStatus `json:"status" binding:"required,oneof=PUBLISHED DRAFT HIDDEN"`
	CreatedAt   time.Time            `json:"createdAt" binding:"required"`
	UpdatedAt   *time.Time           `json:"updatedAt" gorm:"autoUpdateTime:false"`
}

func (a ArticleUpdateDto) TagsModel() []models.Tag {
	tags := make([]models.Tag, len(a.Tags))
	for i, tag := range a.Tags {
		tags[i] = models.Tag{
			Name: tag.Name,
		}
	}
	return tags
}

func ToArticleDto(article models.Article) ArticleDto {
	articleDto := ArticleDto{
		Aid:         article.Aid,
		User:        toArticleUserDto(article.User),
		Category:    toArticleCategoryDto(article.Category),
		Tags:        toArticleTagDtoList(article.Tags),
		Title:       article.Title,
		Abstract:    article.Abstract,
		Content:     article.Content,
		Commentable: article.Commentable,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		Status:      article.Status,
		ReadCount:   article.ReadCount,
	}
	return articleDto
}

func ToArticlePreviewDtoList(articles models.Page[models.Article]) models.Page[ArticlePreviewDto] {
	articleListDto := models.Page[ArticlePreviewDto]{
		Total: articles.Total,
		Page:  articles.Page,
		Size:  articles.Size,
		Data:  make([]ArticlePreviewDto, len(articles.Data)),
	}
	for i, article := range articles.Data {
		articleListDto.Data[i] = ToArticlePreviewDto(article)
	}
	return articleListDto
}

func ToArticlePreviewDto(article models.Article) ArticlePreviewDto {
	articleDto := ArticlePreviewDto{
		Aid:       article.Aid,
		User:      toArticleUserDto(article.User),
		Category:  toArticleCategoryDto(article.Category),
		Tags:      toArticleTagDtoList(article.Tags),
		Title:     article.Title,
		Abstract:  article.Abstract,
		CreatedAt: article.CreatedAt,
		Status:    article.Status,
		ReadCount: article.ReadCount,
	}
	return articleDto
}

func toArticleUserDto(user models.User) ArticleUserDto {
	return ArticleUserDto{
		Uid:  user.Uid,
		Name: user.Name,
	}
}

func toArticleCategoryDto(category models.Category) ArticleCategoryDto {
	return ArticleCategoryDto{
		Cid:  category.Cid,
		Name: category.Name,
	}
}

func toArticleTagDtoList(tags []models.Tag) []ArticleTagDto {
	tagsDto := make([]ArticleTagDto, len(tags))
	for i, tag := range tags {
		tagsDto[i] = ArticleTagDto{
			Name: tag.Name,
		}
	}
	return tagsDto
}
