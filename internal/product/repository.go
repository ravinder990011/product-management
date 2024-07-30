package product

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Repository interface {
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id uint) error
	GetProductByID(id uint) (*Product, error)
	GetProductByName(name string) (*Product, error)
	GetProducts() ([]*Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateProduct(product *Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return errors.Wrap(err, "failed to create product")
	}
	return nil
}

func (r *repository) UpdateProduct(product *Product) error {
	if err := r.db.Save(product).Error; err != nil {
		return errors.Wrap(err, "failed to update product")
	}
	return nil
}

func (r *repository) DeleteProduct(id uint) error {
	if err := r.db.Delete(&Product{}, id).Error; err != nil {
		return errors.Wrap(err, "failed to delete product")
	}
	return nil
}

func (r *repository) GetProductByID(id uint) (*Product, error) {
	var product Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get product by ID")
	}
	return &product, nil
}

func (r *repository) GetProductByName(name string) (*Product, error) {
	var product Product
	if err := r.db.Where("name = ?", name).First(&product).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get product by name")
	}
	return &product, nil
}

func (r *repository) GetProducts() ([]*Product, error) {
	var products []*Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get products")
	}
	return products, nil
}
