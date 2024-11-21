package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPostById(c *gin.Context) {
	ctx := c.Request.Context()

	postIDStr := c.Param("post_id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errors.New("invalid post id").Error(),
		})
		return
	}

	response, err := h.postSvc.GetPostById(ctx, postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
