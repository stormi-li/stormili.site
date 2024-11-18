package controllers

import (
	"comment/database"
	"comment/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 新增评论
func AddComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// 添加创建时间
	comment.CreatedAt = time.Now()

	// 保存评论
	result := database.DB.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}
	
	c.JSON(http.StatusOK, comment)
}

// 查询评论（支持按类别筛选）
func GetComments(c *gin.Context) {
	var comments []models.Comment
	category := c.Query("category")

	query := database.DB.Order("created_at desc")
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// 查询评论
	result := query.Find(&comments)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
