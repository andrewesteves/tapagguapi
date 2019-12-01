package transformer

import "github.com/andrewesteves/tapagguapi/model"

// CategoryTransformer struct
type CategoryTransformer struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

// TransformOne category specified JSON
func (rf CategoryTransformer) TransformOne(category model.Category) CategoryTransformer {
	var newCategory CategoryTransformer
	newCategory.ID = category.ID
	newCategory.Title = category.Title
	newCategory.Icon = category.Icon
	return newCategory
}
