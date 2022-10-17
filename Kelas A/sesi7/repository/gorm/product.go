package gorm

import (
	"sesi7/model"
	"sesi7/repository"

	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) repository.ProductRepository {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) CreateProduct(product *model.Product) error {
	return p.db.Create(product).Error
}

func (p *productRepo) GetProducts() (*[]model.Product, error) {
	var products []model.Product

	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return &products, nil
}
