package dto

// IdParam is a struct for binding id from uri.
type IdParam struct {
	Id uint64 `uri:"id" binding:"required"`
}

// StringParam is a struct for binding string from uri.
type StringParam struct {
	Name string `uri:"name" binding:"required"`
}

// PageParam is a struct for binding page and size from query.
// should be used with GetPage and GetSize.
type PageParam struct {
	Page int `form:"page" binding:"omitempty,gte=1"`
	Size int `form:"size" binding:"omitempty,gte=10,lte=50"`
}

func (p PageParam) GetPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

func (p PageParam) GetSize() int {
	if p.Size == 0 {
		return 10
	}
	return p.Size
}
