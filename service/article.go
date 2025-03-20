package service

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/cache"
	"blackhole-blog/pkg/dao"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/upload"
	"blackhole-blog/pkg/util"
	"errors"
	"fmt"
	"io"
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
	defer cache.DeferredSetWithRecover(cache.Article, cacheKey, &res)()

	article, daoErr := dao.Article.FindById(id)
	panicNotFoundErrIfNotNil(daoErr, "未找到该文章")
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

func (articleService) UploadAttachment(attachment io.Reader, filename string) string {
	path, err := upload.Uploader.UploadFile(attachment, filename)
	if err != nil {
		if errors.Is(err, upload.ErrFileAlreadyExist) {
			panic(util.NewError(http.StatusBadRequest, "文件已存在，请更换文件名"))
		}
		panicErrIfNotNil(err)
	}
	return path
}

func (articleService) Update(article dto.ArticleUpdateDto) {
	// cache
	cacheKey := fmt.Sprintf("article:%d", article.Aid)
	defer cache.Article.Delete(cacheKey)

	err := dao.Article.Update(article)
	panicNotFoundErrIfNotNil(err, "未找到该文章", entryErrProducer(1452, foreignKeyErrProducer))
}

func foreignKeyErrProducer(msg string) string {
	if strings.Contains(msg, "uid") {
		return "未找到该用户"
	}
	if strings.Contains(msg, "cid") {
		return "未找到该分类"
	}
	if strings.Contains(msg, "aid") {
		return "未找到该文章"
	}
	return setting.InternalErrorMessage
}

func (articleService) Delete(id uint64) {
	// cache
	cacheKey := fmt.Sprintf("article:%d", id)
	defer cache.Article.Delete(cacheKey)

	affects, err := dao.Article.Delete(id)
	panicErrIfNotNil(err)
	if affects == 0 {
		panic(util.NewError(http.StatusBadRequest, "未找到该文章"))
	}
}

const abstractPrompt = "你是一个文章总结专家，擅长对文章尤其是技术类的文章提取摘要，你拥有强大的内容分析能力，能准确提取关键信息和核心要点。具备将用户输入的markdown文章提炼成整洁语句的能力，以便让读者快速了解文章内容，你的回答应直接给出结果（即文章摘要纯文本），不需要多余的回复，并尽量控制在100字以内。"

func (articleService) GenerateAbstract(content string, closeNotify <-chan bool) <-chan string {
	return AiChat.StreamingChat(abstractPrompt, content, closeNotify)
}
