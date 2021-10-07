package resp

type Category struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	IsHidden         bool   `json:"isHidden"`
	ParentCategoryId int64  `json:"parentCategoryId"`
	Level            int    `json:"level"`
}
