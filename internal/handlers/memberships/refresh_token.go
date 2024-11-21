package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/situs-forum/internal/models/memberships"
)

func (h *Handler) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	var request memberships.RefreshTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userId")
	accessToken, err := h.membershipSvc.ValidateRefreshToken(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, memberships.RefreshTokenResponse{
		AccessToken: accessToken,
	})

}
