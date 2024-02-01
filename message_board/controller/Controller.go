package controller

import (
	"message/model"
	"message/repository"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

// GetAll
// @Summary 查询全部留言
// @Produce json
// @Success 200 {object} string "成功"
// @Failure 400 {object} string "查询错误"
// @Router /api/v1/message [get]
func GetAll(c *gin.Context) {
	message, err := repository.GetAllMessage()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// @Summary 查询 {id} 留言
// @Produce json
// @Param id path int true "留言ID"
// @Success 200 {object} string "成功"
// @Failure 404 {object} string "找不到留言"
// @Router /api/v1/message/{id} [get]
func Get(c *gin.Context) {
	var message model.Message

	if err := repository.GetMessage(&message, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "找不到留言"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// @Summary 新增留言
// @Produce json
// @Param Content formData string true "留言內容"
// @Success 201 {object} string "成功"
// @Failure 400 {object} string "新增留言錯誤"
// @Router /api/v1/message [post]
func Create(c *gin.Context) {
	var message model.Message

	if c.PostForm("Content") == "" || utf8.RuneCountInString(c.PostForm("Content")) >= 20 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "沒有輸入內容或長度超過20個字元"})
		return
	}

	c.Bind(&message)
	repository.CreateMessage(&message)
	c.JSON(http.StatusCreated, gin.H{"message": message})
}

// @Summary 更新留言
// @Produce json
// @Param id path int true "留言ID"
// @Param Content formData string true "留言內容"
// @Success 200 {object} string "成功"
// @Failure 400 {object} string "更新留言錯誤"
// @Router /api/v1/message/{id} [put]
func Update(c *gin.Context) {
	var message model.Message

	if c.PostForm("Content") == "" || utf8.RuneCountInString(c.PostForm("Content")) >= 20 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "沒有輸入內容或長度超過20個字元"})
		return
	}

	if err := repository.UpdateMessage(&message, c.PostForm("Content"), c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "找不到留言"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// @Summary 刪除留言
// @Produce json
// @Param id path int true "留言ID"
// @Success 204 {object} string "成功"
// @Failure 404 {object} string "找不到留言"
// @Router /api/v1/message/{id} [delete]
func Delete(c *gin.Context) {
	var message model.Message

	if err := repository.DeleteMessage(&message, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "找不到留言"})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "刪除留言成功"})
}
