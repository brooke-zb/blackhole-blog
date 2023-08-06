package service

import (
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/dao"
)

type categoryService struct{}

func (categoryService) FindList() []dto.CategoryHeatDto {
	categories, daoErr := dao.Category.FindListWithArticleCount()
	panicErrIfNotNil(daoErr)
	return dto.ToCategoryHeatDtoList(categories)
}
