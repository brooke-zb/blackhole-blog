package service

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/cache"
	"blackhole-blog/pkg/dao"
)

type categoryService struct{}

func (categoryService) FindById(id uint64) dto.CategoryDto {
	category, daoErr := dao.Category.FindById(id)
	panicSelectErrIfNotNil(daoErr, "分类不存在")
	return dto.ToCategoryDto(category)
}

func (categoryService) FindList(page, size int) models.Page[dto.CategoryDto] {
	categories, daoErr := dao.Category.FindList(page, size)
	panicErrIfNotNil(daoErr)
	return dto.ToCategoryDtoList(categories)
}

func (categoryService) FindListWithArticleCount() []dto.CategoryHeatDto {
	categories, daoErr := dao.Category.FindListWithArticleCount()
	panicErrIfNotNil(daoErr)
	return dto.ToCategoryHeatDtoList(categories)
}

func (categoryService) Add(category dto.CategoryAddDto) {
	daoErr := dao.Category.Add(category)
	panicErrIfNotNil(daoErr, entryErr(1062, "分类名已存在"))
}

func (categoryService) Update(category dto.CategoryUpdateDto) {
	// cache
	defer cache.Article.Clear()

	daoErr := dao.Category.Update(category)
	panicErrIfNotNil(daoErr, entryErr(1062, "分类名已存在"))
}

func (categoryService) Delete(id uint64) {
	// cache
	defer cache.Article.Clear()

	daoErr := dao.Category.Delete(id)
	panicErrIfNotNil(daoErr, entryErr(1451, "该分类下存在文章，禁止删除"))
}
