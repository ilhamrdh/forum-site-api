package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/situs-forum/internal/models/posts"
)

func (h *Handler) CreateCommnet(c *gin.Context) {
	var request posts.CreateCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	postIDStr := c.Param("post_id")
	postIDInt, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errors.New("post id tidak valid").Error(),
		})
		return
	}

	userID := c.GetInt64("userId")

	err = h.postSvc.CreateCommnet(c.Request.Context(), postIDInt, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)

}
