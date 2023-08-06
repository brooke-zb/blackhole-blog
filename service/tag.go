package service

import (
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/dao"
)

type tagService struct{}

func (tagService) FindListWithHeat() []dto.TagHeatDto {
	tags, daoErr := dao.Tag.FindListWithHeat()
	panicErrIfNotNil(daoErr)
	return dto.ToTagHeatDtoList(tags)
}
