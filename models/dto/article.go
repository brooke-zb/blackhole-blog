package dto

import (
	"blackhole-blog/models"
	"time"
)

type ArticleDto struct {
	Aid         uint64             `json:"aid"`
	Uid         uint64             `json:"uid"`
	Category    ArticleCategoryDto `json:"category"`
	Tags        []ArticleTagDto    `json:"tags"`
	Title       string             `json:"title"`
	Content     string             `json:"content"`
	Commentable bool               `json:"commentable"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   *time.Time         `json:"UpdatedAt"`
	Status      string             `json:"status"`
	ReadCount   int                `json:"readCount"`
}

type ArticlePreviewDto struct {
	Aid       uint64             `json:"aid"`
	Category  ArticleCategoryDto `json:"category"`
	Tags      []ArticleTagDto    `json:"tags"`
	Title     string             `json:"title"`
	CreatedAt time.Time          `json:"createdAt"`
	Status    string             `json:"status"`
	ReadCount int                `json:"readCount"`
}

type ArticleCategoryDto struct {
	Cid  uint64 `json:"cid"`
	Name string `json:"name"`
}

type ArticleTagDto struct {
	Name string `json:"name"`
}

func ToArticleDto(article models.Article) ArticleDto {
	articleDto := ArticleDto{
		Aid: article.Aid,
		Uid: article.Uid,
		Category: ArticleCategoryDto{
			Cid:  article.Category.Cid,
			Name: article.Category.Name,
		},
		Tags:        make([]ArticleTagDto, len(article.Tags)),
		Title:       article.Title,
		Content:     article.Content,
		Commentable: article.Commentable,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		Status:      article.Status,
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
			Aid: article.Aid,
			Category: ArticleCategoryDto{
				Cid:  article.Category.Cid,
				Name: article.Category.Name,
			},
			Tags:      make([]ArticleTagDto, len(article.Tags)),
			Title:     article.Title,
			CreatedAt: article.CreatedAt,
			Status:    article.Status,
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
