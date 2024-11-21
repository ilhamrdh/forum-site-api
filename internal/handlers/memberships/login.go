package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/situs-forum/internal/models/memberships"
	"github.com/rs/zerolog/log"
)

func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var request memberships.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, refreshToken, err := h.membershipSvc.Login(ctx, request)
	log.Print("Handler -> refresh token:", refreshToken)
	log.Print("Handler -> token:", token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := memberships.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, response)
}
