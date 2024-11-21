package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/situs-forum/internal/models/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	var request memberships.SignUpRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.membershipSvc.SignUp(ctx, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register successfully",
	})
}
