package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	ID               int64     `gorm:"primary_key" json:"id"`
	Name             string    `json:"name"`
	IsHidden         bool      `json:"isHidden"`
	ParentCategoryId int64     `json:"parentCategoryId"`
	Level            int       `json:"level"`
	Created          time.Time `json:"created"`
}

func GetCategoryWithId(id int) (Category, error) {
	var category Category
	result := db.Where("id = ?", id).Where("is_hidden = ?", false).First(&category)
	if result.Error != nil {
		return category, result.Error
	}
	return category, nil
}

func AddCategory(name string, parentId int64) bool {
	db.Create(&Category{
		Name:             name,
		ParentCategoryId: parentId,
	})

	return true
}

func (category *Category) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Created", time.Now())

	return nil
}
