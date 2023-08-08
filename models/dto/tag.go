package dto

import "blackhole-blog/models"

type TagHeatDto struct {
	Tid          uint64 `json:"tid"`
	Name         string `json:"name"`
	ArticleCount int    `json:"articleCount"`
}

type TagDto struct {
	Tid  uint64 `json:"tid" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type TagAddDto struct {
	Name string `json:"name" binding:"required"`
}

type TagUpdateDto TagDto

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

func ToTagDtoList(tags models.Page[models.Tag]) models.Page[TagDto] {
	tagDtoList := models.Page[TagDto]{
		Page:  tags.Page,
		Size:  tags.Size,
		Total: tags.Total,
		Data:  make([]TagDto, len(tags.Data)),
	}
	for i, tag := range tags.Data {
		tagDtoList.Data[i] = ToTagDto(tag)
	}
	return tagDtoList
}

func ToTagDto(tag models.Tag) TagDto {
	return TagDto{
		Tid:  tag.Tid,
		Name: tag.Name,
	}
}
