package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"store/models"
	"strconv"
)

type getProductResponse struct {
	Data  []models.Product `json:"data"`
	Page  int              `json:"page"`
	Total int64            `json:"total"`
}

func (h *Handler) createProduct(c *gin.Context) {
	var input models.Product

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Product.Create(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getProduct(c *gin.Context) {
	page, ok := c.GetQuery("page")
	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "no page provided")
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "page type conversion error")
		return
	}

	limit, ok := c.GetQuery("limit")
	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "no offset provided")
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "offset type conversion error")
		return
	}

	product, total := h.services.Product.Get(pageInt, limitInt)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getProductResponse{Data: product, Page: pageInt, Total: total})
}

func (h *Handler) getLatestProduct(c *gin.Context) {
	product := h.services.Product.Latest()
	c.JSON(http.StatusOK, getProductResponse{Data: product})
}

func (h *Handler) deleteProduct(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "id type conversion error")
	}
	h.services.Product.Delete(uint(ID))
}
