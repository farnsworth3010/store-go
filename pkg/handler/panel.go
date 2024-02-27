package handler

import (
	"net/http"
	"store/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAdmins(c *gin.Context) {
	res, err := h.services.Panel.GetAdmins()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func (h *Handler) getUsers(c *gin.Context) {
	res, err := h.services.Panel.GetUsers()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func (h *Handler) getBlogs(c *gin.Context) {
	res, err := h.services.Panel.GetBlogs()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func (h *Handler) setRole(c *gin.Context) {
	var input models.SetRoleInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Panel.SetRole(uint(input.ID), uint(input.RoleID))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
