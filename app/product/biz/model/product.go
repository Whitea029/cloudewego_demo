package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float32    `json:"price"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
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
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	return
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{ctx: ctx, db: db}
}

type CacheProductQuery struct {
	productQuery ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func (c CacheProductQuery) GetById(productID uint) (product Product, err error) {
	key := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productID)
	cacheRes := c.cacheClient.Get(c.productQuery.ctx, key)
	err = func() error {
		if err := cacheRes.Err(); err != nil {
			return err
		}
		cacheResByte, err := cacheRes.Bytes()
		if err != nil {
			return err
		}
		err = json.Unmarshal(cacheResByte, &product)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		product, err = c.productQuery.GetById(productID)
		if err != nil {
			return Product{}, err
		}
		productByte, err := json.Marshal(product)
		if err != nil {
			return product, err
		}
		if err := c.cacheClient.Set(c.productQuery.ctx, key, productByte, time.Hour).Err(); err != nil {
			return product, err
		}
	}
	return
}

func (c CacheProductQuery) SearchProduct(q string) (products []Product, err error) {
	err = c.productQuery.db.WithContext(c.productQuery.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?", "%"+q+"%", "%"+q+"%").Error
	return
}

func NewCacheProductQuery(ctx context.Context, db *gorm.DB, cacheClient *redis.Client) *CacheProductQuery {
	return &CacheProductQuery{productQuery: *NewProductQuery(ctx, db), cacheClient: cacheClient, prefix: "whmall"}
}
