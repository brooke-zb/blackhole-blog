package dto

import "blackhole-blog/models"

type CategoryHeatDto struct {
	Cid          uint64 `json:"cid"`
	Name         string `json:"name"`
	ArticleCount int    `json:"articleCount"`
}

type CategoryDto struct {
	Cid  uint64 `json:"cid" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CategoryAddDto struct {
	Name string `json:"name" binding:"required"`
}

type CategoryUpdateDto CategoryDto

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

func ToCategoryDtoList(categories models.Page[models.Category]) models.Page[CategoryDto] {
	categoryDtoList := models.Page[CategoryDto]{
		Page:  categories.Page,
		Size:  categories.Size,
		Total: categories.Total,
		Data:  make([]CategoryDto, len(categories.Data)),
	}
	for i, category := range categories.Data {
		categoryDtoList.Data[i] = ToCategoryDto(category)
	}
	return categoryDtoList
}

func ToCategoryDto(category models.Category) CategoryDto {
	return CategoryDto{
		Cid:  category.Cid,
		Name: category.Name,
	}
}
