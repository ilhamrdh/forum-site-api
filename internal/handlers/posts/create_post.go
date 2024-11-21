package posts

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/situs-forum/internal/models/posts"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	userId := c.GetInt64("userId")
	log.Println("User ID", userId)

	err := h.postSvc.CreatePost(ctx, userId, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success create posts",
	})
}
