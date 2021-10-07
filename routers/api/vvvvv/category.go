package vvvvv

import (
	"net/http"
	"nueip/eshop/models"
	"nueip/eshop/pkg/e"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	Name     string `json:"name" binding:"required,len=10"`
	ParentId int64  `json:"parentId" binding:""`
}

func GetCategory(c *gin.Context) {

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
		category, err := models.GetCategoryWithId(id)
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
			"data": category,
		})
		return

	}

}

func CreateCategory(c *gin.Context) {
	var req createCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {

	}
	if models.AddCategory(req.Name, req.ParentId) {
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  e.GetMsg(e.SUCCESS),
			"data": make(map[string]string),
		})
	}

}
func EditCategory(c *gin.Context) {

}
func DeleteCategory(c *gin.Context) {

}
