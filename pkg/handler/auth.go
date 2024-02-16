package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"store/models"
)

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserInfoInput struct {
	Token string `json:"token" binding:"required"`
}

// @Summary sign in
// @Schemes
// @Description sign in
// @Tags auth
// @Param request body handler.signInInput true "query params"
// @Accept json
// @Produce json
// @Success 200 {string} string "token"
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}

// @Summary sign up
// @Schemes
// @Description sign up
// @Tags auth
// @Param request body models.User true "query params"
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getUserInfo(c *gin.Context) {
	var input GetUserInfoInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.Authorization.GetUserInfo(input.Token)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
