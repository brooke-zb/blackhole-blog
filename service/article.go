package service

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/cache"
	"blackhole-blog/pkg/dao"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"fmt"
	"net/http"
	"strings"
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

func (articleService) FindList(clause models.ArticleClause) (res models.Page[dto.ArticlePreviewDto]) {
	articles, daoErr := dao.Article.FindPreviewList(clause)
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

func (articleService) Add(article dto.ArticleAddDto) {
	a := article.ToArticleModel()
	a.Aid = util.NextId()
	err := dao.Article.Add(a)
	panicErrIfNotNil(err, entryErrProducer(1452, foreignKeyErrProducer))
}

func (articleService) Update(article dto.ArticleUpdateDto) {
	err := dao.Article.Update(article)
	panicErrIfNotNil(err, entryErrProducer(1452, foreignKeyErrProducer))
}

func foreignKeyErrProducer(msg string) string {
	if strings.Contains(msg, "uid") {
		return "用户不存在"
	}
	if strings.Contains(msg, "cid") {
		return "分类不存在"
	}
	if strings.Contains(msg, "aid") {
		return "文章不存在"
	}
	return setting.InternalErrorMessage
}

func (articleService) Delete(id uint64) {
	affects, err := dao.Article.Delete(id)
	panicErrIfNotNil(err)
	if affects == 0 {
		panic(util.NewError(http.StatusBadRequest, "文章不存在"))
	}
}
