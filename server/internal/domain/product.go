package domain

import (
	"errors"
	"time"
)

type ProductID uint64
type ProductPrice uint64

type ProductPageSize uint16

type Product struct {
	ID          uint64    `json:"id" form:"id" xml:"id"`
	Name        string    `json:"name" form:"name" xml:"form"`
	Description string    `json:"description" form:"description" xml:"description"`
	Price       uint64    `json:"price" form:"price" xml:"price"`
	CreatedAt   time.Time `json:"created_at" form:"created_at" xml:"created_at"`
}

type ProductInput struct {
	Name        string `json:"name" form:"name" xml:"id" binding:"required"`
	Description string `json:"description" form:"description" xml:"description" binding:"required"`
	Price       uint64 `json:"price" form:"price" xml:"price" binding:"required"`
}

var (
	ErrProductNotFound   = errors.New("Product not found")
	MaxProductNameLength = 255
)
