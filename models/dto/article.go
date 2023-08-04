package dto

import (
	"blackhole-blog/models"
	"time"
)

type ArticleDto struct {
	Aid         uint64
	Category    ArticleCategoryDto
	Tags        []ArticleTagDto
	Title       string
	Content     string
	Commentable bool
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	ReadCount   int
}

type ArticlePreviewDto struct {
	Aid       uint64
	Category  ArticleCategoryDto
	Tags      []ArticleTagDto
	Title     string
	CreatedAt time.Time
	ReadCount int
}

type ArticleCategoryDto = models.Category

type ArticleTagDto struct {
	Name string
}

func ToArticleDto(article models.Article) ArticleDto {
	articleDto := ArticleDto{
		Aid:         article.Aid,
		Category:    article.Category,
		Tags:        make([]ArticleTagDto, len(article.Tags)),
		Title:       article.Title,
		Content:     article.Content,
		Commentable: article.Commentable,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		ReadCount:   article.ReadCount,
	}
	for i, tag := range article.Tags {
		articleDto.Tags[i] = ArticleTagDto{
			Name: tag.Name,
		}
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
		articleListDto.Data[i] = ArticlePreviewDto{
			Aid:       article.Aid,
			Category:  article.Category,
			Tags:      make([]ArticleTagDto, len(article.Tags)),
			Title:     article.Title,
			CreatedAt: article.CreatedAt,
			ReadCount: article.ReadCount,
		}
		for j, tag := range article.Tags {
			articleListDto.Data[i].Tags[j] = ArticleTagDto{
				Name: tag.Name,
			}
		}
	}
	return articleListDto
}
