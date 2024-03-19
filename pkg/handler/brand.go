package handler

import (
	"net/http"
	"store/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getBrandResponse struct {
	Data  []models.Brand `json:"data"`
	Total int64          `json:"total"`
}

// @Summary Create brand post
// @Schemes
// @Description creates post
// @Tags brand
// @Param request body models.CreateBrandParams true "query params"
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Router /brand [post]
func (h *Handler) createBrand(c *gin.Context) {
	var input models.CreateBrandParams

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Brand.Create(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get brand post
// @Schemes
// @Description gets posts
// @Tags brand
// @Param limit   query int true "limit"
// @Param page query int true "page"
// @Produce json
// @Success 200 {object} handler.getBrandResponse
// @Router /brand [get]
func (h *Handler) getBrand(c *gin.Context) {
	brand, total := h.services.Brand.Get()
	c.JSON(http.StatusOK, getBrandResponse{Data: brand, Total: total})
}

func (h *Handler) deleteBrand(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "id type conversion error")
	}
	h.services.Brand.Delete(uint(ID))
}

func (h *Handler) updateBrand(c *gin.Context) {
	var input models.EditBrandParams

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Brand.Update(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
