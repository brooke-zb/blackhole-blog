package service

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/cache"
	"blackhole-blog/pkg/dao"
)

type tagService struct{}

func (tagService) FindById(id uint64) dto.TagDto {
	tag, daoErr := dao.Tag.FindById(id)
	panicNotFoundErrIfNotNil(daoErr, "未找到该标签")
	return dto.ToTagDto(tag)
}

func (tagService) FindList(page, size int) models.Page[dto.TagDto] {
	tags, daoErr := dao.Tag.FindList(page, size)
	panicErrIfNotNil(daoErr)
	return dto.ToTagDtoList(tags)
}

func (tagService) FindListWithHeat() []dto.TagHeatDto {
	tags, daoErr := dao.Tag.FindListWithHeat()
	panicErrIfNotNil(daoErr)
	return dto.ToTagHeatDtoList(tags)
}

func (tagService) Add(tag dto.TagAddDto) {
	daoErr := dao.Tag.Add(tag)
	panicErrIfNotNil(daoErr, entryErr(1062, "标签名已存在"))
}

func (tagService) Update(tag dto.TagUpdateDto) {
	// cache
	defer cache.Article.Clear()

	daoErr := dao.Tag.Update(tag)
	panicNotFoundErrIfNotNil(daoErr, "未找到该标签", entryErr(1062, "标签名已存在"))
}

func (tagService) DeleteBatch(ids ...uint64) int64 {
	// cache
	defer cache.Article.Clear()

	affects, daoErr := dao.Tag.DeleteBatch(ids...)
	panicErrIfNotNil(daoErr)
	return affects
}
