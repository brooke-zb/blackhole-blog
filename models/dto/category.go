package dto

import "blackhole-blog/models"

type CategoryHeatDto struct {
	Cid          uint64 `json:"cid"`
	Name         string `json:"name"`
	ArticleCount int    `json:"articleCount"`
}

func ToCategoryHeatDtoList(categories []models.Category) []CategoryHeatDto {
	categoryHeatDtoList := make([]CategoryHeatDto, len(categories))
	for i, category := range categories {
		categoryHeatDtoList[i] = ToCategoryHeatDto(category)
	}
	return categoryHeatDtoList
}

func ToCategoryHeatDto(category models.Category) CategoryHeatDto {
	return CategoryHeatDto{
		Cid:          category.Cid,
		Name:         category.Name,
		ArticleCount: category.ArticleCount,
	}
}
