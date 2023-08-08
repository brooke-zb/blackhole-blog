package service

import (
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/dao"
)

type tagService struct{}

func (tagService) FindById(id uint64) dto.TagDto {
	tag, daoErr := dao.Tag.FindById(id)
	panicSelectErrIfNotNil(daoErr, "标签不存在")
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
	daoErr := dao.Tag.Update(tag)
	panicErrIfNotNil(daoErr, entryErr(1062, "标签名已存在"))
}

func (tagService) DeleteBatch(ids ...uint64) int64 {
	affects, daoErr := dao.Tag.DeleteBatch(ids...)
	panicErrIfNotNil(daoErr)
	return affects
}
