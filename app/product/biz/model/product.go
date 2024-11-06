package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_category;"`
}

func (Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(productID uint) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productID).Error
	return
}

func (p ProductQuery) SearchProduct(q string) (products []Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like or description like ?", "%"+q+"%", "%"+q+"%").Error
	return
}
