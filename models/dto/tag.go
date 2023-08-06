package dto

import "blackhole-blog/models"

type TagHeatDto struct {
	Tid          uint64 `json:"tid"`
	Name         string `json:"name"`
	ArticleCount int    `json:"articleCount"`
}

func ToTagHeatDtoList(tags []models.Tag) []TagHeatDto {
	tagHeatDtoList := make([]TagHeatDto, len(tags))
	for i, tag := range tags {
		tagHeatDtoList[i] = ToTagHeatDto(tag)
	}
	return tagHeatDtoList
}

func ToTagHeatDto(tag models.Tag) TagHeatDto {
	return TagHeatDto{
		Tid:          tag.Tid,
		Name:         tag.Name,
		ArticleCount: tag.ArticleCount,
	}
}
