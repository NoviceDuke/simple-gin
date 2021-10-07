package controller

import (
	"fmt"
	"net/http"
	"nueip/eshop/models"
	"nueip/eshop/pkg/e"
	"nueip/eshop/resp"
	"nueip/eshop/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	Name     string `json:"name" `
	ParentId int64  `json:"parentId" `
}
type CategoryController struct {
	CategorySrv service.CategoryService
}

func (ctr *CategoryController) GetCategory(c *gin.Context) {
	if len(c.Param("id")) > 0 {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"code": e.ERROR,
				"msg":  e.GetMsg(e.INVALID_PARAMS),
				"data": make(map[string]string),
			})
			return
		}
		category := ctr.CategorySrv.ExistByCategoryID(id)
		if err != nil || category.IsHidden {
			c.JSON(http.StatusNotFound, gin.H{
				"code": e.ERROR_NOT_EXIST_CATEGORY,
				"msg":  e.GetMsg(e.ERROR_NOT_EXIST_CATEGORY),
				"data": make(map[string]string),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  e.GetMsg(e.SUCCESS),
			"data": resp.Category{
				Id:               int(category.ID),
				Name:             category.Name,
				IsHidden:         category.IsHidden,
				ParentCategoryId: category.ParentCategoryId,
			},
		})

	}
}

// func (ctr *CategoryController) CategoryInfo(c *gin.Context) {
// 	entity := resp.Entity{
// 		Code:      int(enum.OperateFail),
// 		Msg:       enum.OperateFail.String(),
// 		Total:     0,
// 		TotalPage: 1,
// 		Data:      nil,
// 	}
// 	userId := c.Param("id")
// 	if userId == "" {
// 		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
// 		return
// 	}
// 	u := model.User{
// 		UserId: userId,
// 	}
// 	result, err := h.UserSrv.Get(u)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
// 		return
// 	}

// 	r := h.GetEntity(*result)

// 	entity = resp.Entity{
// 		Code:      http.StatusOK,
// 		Msg:       "OK",
// 		Total:     0,
// 		TotalPage: 0,
// 		Data:      r,
// 	}
// 	c.JSON(http.StatusOK, gin.H{"entity": entity})
// }

// func (ctr *CategoryController) UserListHandler(c *gin.Context) {
// 	var q query.ListQuery
// 	entity := resp.Entity{
// 		Code:      int(enum.OperateFail),
// 		Msg:       enum.OperateFail.String(),
// 		Total:     0,
// 		TotalPage: 1,
// 		Data:      nil,
// 	}
// 	err := c.ShouldBindQuery(&q)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
// 		return
// 	}
// 	list, err := h.UserSrv.List(&q)
// 	total, err := h.UserSrv.GetTotal(&q)

// 	if err != nil {
// 		panic(err)
// 	}
// 	if q.PageSize == 0 {
// 		q.PageSize = 5
// 	}
// 	ret := int(total) % q.PageSize
// 	ret2 := int(total) / q.PageSize
// 	totalPage := 0
// 	if ret == 0 {
// 		totalPage = ret2
// 	} else {
// 		totalPage = ret2 + 1
// 	}
// 	var newList []*resp.User
// 	for _, item := range list {
// 		r := h.GetEntity(*item)
// 		newList = append(newList, &r)
// 	}

// 	entity = resp.Entity{
// 		Code:      http.StatusOK,
// 		Msg:       "OK",
// 		Total:     total,
// 		TotalPage: totalPage,
// 		Data:      newList,
// 	}
// 	c.JSON(http.StatusOK, gin.H{"entity": entity})
// }

func (ctr *CategoryController) CreateCategory(c *gin.Context) {

	u := models.Category{}
	err := c.ShouldBindJSON(&u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.ERROR,
			"msg":  e.GetMsg(e.INVALID_PARAMS),
			"data": make(map[string]string),
		})
		return
	}

	_, err = ctr.CategorySrv.Add(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": e.ERROR,
			"msg":  e.GetMsg(e.ERROR),
			"data": make(map[string]string),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": make(map[string]string),
	})

}

func (ctr *CategoryController) EditCategory(c *gin.Context) {
	if len(c.Param("id")) > 0 {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"code": e.ERROR,
				"msg":  e.GetMsg(e.INVALID_PARAMS),
				"data": make(map[string]string),
			})
			return
		}

		u := models.Category{
			ID: int64(id),
		}

		err = c.ShouldBindJSON(&u)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": e.ERROR,
				"msg":  e.GetMsg(e.INVALID_PARAMS),
				"data": make(map[string]string),
			})
			return
		}
		fmt.Println("aaaaa")
		done, err := ctr.CategorySrv.Edit(u)
		if err != nil || !done {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": e.ERROR,
				"msg":  e.GetMsg(e.ERROR),
				"data": make(map[string]string),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  e.GetMsg(e.SUCCESS),
			"data": make(map[string]string),
		})
	}
}

// func (ctr *CategoryController) DeleteUserHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	b, err := h.UserSrv.Delete(id)
// 	entity := resp.Entity{
// 		Code:  int(enum.OperateFail),
// 		Msg:   enum.OperateFail.String(),
// 		Total: 0,
// 		Data:  nil,
// 	}
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{"entity": entity})
// 		return
// 	}
// 	if b {
// 		entity.Code = int(enum.OperateOk)
// 		entity.Msg = enum.OperateOk.String()
// 		c.JSON(http.StatusOK, gin.H{"entity": entity})
// 	}
// }
