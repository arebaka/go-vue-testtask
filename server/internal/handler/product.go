package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"are.moe/testtask/internal/config"
	"are.moe/testtask/internal/domain"
	"are.moe/testtask/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	AddProduct(ctx *gin.Context)
	GetProductByID(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	RemoveProduct(ctx *gin.Context)
	ListProducts(ctx *gin.Context)
	FindProductsByName(ctx *gin.Context)
}

type Product struct {
	service service.ProductService
}

func parseOffsetAndLimit(ctx *gin.Context) (domain.ProductID, domain.ProductPageSize, error) {
	offset, err := strconv.ParseUint(ctx.DefaultQuery("offset", "0"), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	limit, err := strconv.ParseUint(ctx.DefaultQuery("limit", "0"), 10, 16)
	if err != nil {
		return 0, 0, nil
	}
	if limit == 0 {
		limit = uint64(config.Get().PageSize)
	}

	return domain.ProductID(offset), domain.ProductPageSize(limit), nil
}

func NewProductHandler(service service.ProductService) *Product {
	return &Product{service}
}

func (h *Product) AddProduct(ctx *gin.Context) {
	var input domain.ProductInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if len(input.Name) > domain.MaxProductNameLength || len(input.Name) < domain.MinProductNameLength {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	product, err := h.service.Create(input)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *Product) GetProductByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("productID"), 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	product, err := h.service.GetByID(domain.ProductID(id))
	if errors.Is(err, domain.ErrProductNotFound) {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *Product) UpdateProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("productID"), 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var input domain.ProductInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if len(input.Name) > domain.MaxProductNameLength || len(input.Name) < domain.MinProductNameLength {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	product, err := h.service.Update(domain.ProductID(id), input)
	if errors.Is(err, domain.ErrProductNotFound) {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *Product) RemoveProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("productID"), 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.service.Delete(domain.ProductID(id))
	if errors.Is(err, domain.ErrProductNotFound) {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Product) ListProducts(ctx *gin.Context) {
	offset, limit, err := parseOffsetAndLimit(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	products, nTotal, err := h.service.List(offset, limit)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, domain.ProductListResponse{Products: products, NTotal: nTotal})
}

func (h *Product) FindProductsByName(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	offset, limit, err := parseOffsetAndLimit(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	products, nTotal, err := h.service.FindByName(name, offset, limit)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, domain.ProductListResponse{Products: products, NTotal: nTotal})
}
