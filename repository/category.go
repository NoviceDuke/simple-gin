package repository

import (
	"fmt"
	"nueip/eshop/models"

	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

type CategoryRepoInterface interface {
	Exist(category models.Category) *models.Category
	ExistByCategoryID(id int) *models.Category
	Add(category models.Category) (*models.Category, error)
	Edit(category models.Category) (bool, error)
}

func (repo *CategoryRepository) Exist(category models.Category) *models.Category {
	var count int
	repo.DB.Find(&category).Where("name = ?", category.Name)
	if count > 0 {
		return &category
	}
	return nil
}

func (repo *CategoryRepository) ExistByCategoryID(id int) *models.Category {
	var category models.Category
	repo.DB.Where("id = ?", id).First(&category)
	return &category
}

func (repo *CategoryRepository) Add(category models.Category) (*models.Category, error) {
	if exist := repo.Exist(category); exist != nil {
		return nil, fmt.Errorf("商品目錄已存在")
	}
	err := repo.DB.Create(&category).Error
	if err != nil {
		return nil, fmt.Errorf("商品目錄建立失敗")
	}
	return &category, nil
}

func (repo *CategoryRepository) Edit(category models.Category) (bool, error) {
	err := repo.DB.Model(&category).Where("id=?", category.ID).Updates(map[string]interface{}{"name": category.Name, "is_hidden": category.IsHidden}).Error
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
