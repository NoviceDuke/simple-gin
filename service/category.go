package service

import (
	"errors"
	"fmt"
	"nueip/eshop/models"
	"nueip/eshop/repository"
)

type CategorySrv interface {
	Exist(category models.Category) *models.Category
	ExistByCategoryID(id string) *models.Category
	Add(category models.Category) (*models.Category, error)
	Edit(category models.Category) (bool, error)
	Delete(id string) (bool, error)
}

type CategoryService struct {
	Repo repository.CategoryRepoInterface
}

func (srv *CategoryService) Add(category models.Category) (*models.Category, error) {
	result := srv.Repo.Exist(category)
	if result != nil {
		return nil, errors.New("商品目錄已經存在")
	}

	return srv.Repo.Add(category)
}
func (srv *CategoryService) Exist(category models.Category) *models.Category {
	return srv.Repo.Exist(category)
}

func (srv *CategoryService) ExistByCategoryID(id int) *models.Category {
	return srv.Repo.ExistByCategoryID(id)
}

func (srv *CategoryService) Edit(category models.Category) (bool, error) {
	if category.ID == 0 {
		return false, fmt.Errorf("參數錯誤")
	}
	exist := srv.Repo.ExistByCategoryID(int(category.ID))
	if exist.ID == 0 {
		return false, errors.New("參數錯誤")
	}
	exist.Name = category.Name
	exist.IsHidden = category.IsHidden
	return srv.Repo.Edit(*exist)
}
