package service

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/cache"
	"blackhole-blog/pkg/dao"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"fmt"
)

type articleService struct{}

func (articleService) FindById(id uint64) (res dto.ArticleDto) {
	// cache
	cacheKey := fmt.Sprintf("article:%d", id)
	articleCache := cache.Article.Get(cacheKey)
	if articleCache != nil && !articleCache.Expired() {
		return articleCache.Value()
	}
	defer cache.DeferredSetWithRevocer(cache.Article, cacheKey, &res)()

	article, daoErr := dao.Article.FindById(id)
	panicSelectErrIfNotNil(daoErr, "未找到该文章")
	return dto.ToArticleDto(article)
}

func (articleService) FindList(page int, size int) (res models.Page[dto.ArticlePreviewDto]) {
	articles, daoErr := dao.Article.FindPreviewList(models.ArticleClause{
		PageParam: models.PageParam{
			PageVal: page,
			SizeVal: size,
		},
	})
	panicErrIfNotNil(daoErr)
	return dto.ToArticlePreviewDtoList(articles)
}

func (articleService) FindListByTag(tag string, page int, size int) (res models.Page[dto.ArticlePreviewDto]) {
	articles, daoErr := dao.Article.FindPreviewList(models.ArticleClause{
		Tag: &tag,
		PageParam: models.PageParam{
			PageVal: page,
			SizeVal: size,
		},
	})
	panicErrIfNotNil(daoErr)
	return dto.ToArticlePreviewDtoList(articles)
}

func (articleService) FindListByCategory(category string, page int, size int) (res models.Page[dto.ArticlePreviewDto]) {
	articles, daoErr := dao.Article.FindPreviewList(models.ArticleClause{
		Category: &category,
		PageParam: models.PageParam{
			PageVal: page,
			SizeVal: size,
		},
	})
	panicErrIfNotNil(daoErr)
	return dto.ToArticlePreviewDtoList(articles)
}

// IncrAndGetReadCount increase and get article read count increment.
func (articleService) IncrAndGetReadCount(id uint64, ip string) int {
	err := util.Redis.PFAdd(getArticleReadCountKey(id), ip)
	panicErrIfNotNil(err)
	count, err := util.Redis.PFCount(getArticleReadCountKey(id))
	panicErrIfNotNil(err)
	return int(count)
}

func getArticleReadCountKey(id uint64) string {
	return fmt.Sprintf("%s%d", setting.ArticleReadCountPrefix, id)
}

func (articleService) UpdateReadCount(aid uint64, incr int64) {
	// cache
	cacheKey := fmt.Sprintf("article:%d", aid)
	defer cache.Article.Delete(cacheKey)

	err := dao.Article.UpdateReadCount(aid, incr)
	panicErrIfNotNil(err)
}
