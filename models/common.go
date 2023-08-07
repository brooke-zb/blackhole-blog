package models

// IdParam is a struct for binding id from uri.
type IdParam struct {
	Id uint64 `uri:"id" binding:"required"`
}

// StringParam is a struct for binding string from uri.
type StringParam struct {
	Name string `uri:"name" binding:"required"`
}

// PageParam is a struct for binding page and size from query.
// should be used with Page and Size.
type PageParam struct {
	PageVal int `form:"page" binding:"omitempty,gte=1"`
	SizeVal int `form:"size" binding:"omitempty,gte=10,lte=50"`
}

func (p PageParam) Page() int {
	if p.PageVal == 0 {
		return 1
	}
	return p.PageVal
}

func (p PageParam) Size() int {
	if p.SizeVal == 0 {
		return 10
	}
	return p.SizeVal
}
