package product

import "github.com/ravinder990011/product-management/util"

type Service interface {
	CreateProduct(name, description string, price float64, stock int) (*Product, error)
	UpdateProduct(id uint, name, description string, price float64, stock int) (*Product, error)
	DeleteProduct(id uint) error
	GetProductByID(id uint, currency string) (*Product, error)
	GetProductByName(name, currency string) (*Product, error)
	GetProducts() ([]*Product, error)
	BulkUpdateProducts(products []*Product) (bool, []uint, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateProduct(name, description string, price float64, stock int) (*Product, error) {
	product := &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}
	if err := s.repo.CreateProduct(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *service) UpdateProduct(id uint, name, description string, price float64, stock int) (*Product, error) {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	if name != "" {
		product.Name = name
	}
	if description != "" {
		product.Description = description
	}
	if price != 0 {
		product.Price = price
	}
	if stock != 0 {
		product.Stock = stock
	}
	if err := s.repo.UpdateProduct(product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *service) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}

func (s *service) GetProductByID(id uint, currency string) (*Product, error) {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	product.Price = util.ConvertCurrency(product.Price, "INR", currency)
	return product, nil
}

func (s *service) GetProductByName(name, currency string) (*Product, error) {
	product, err := s.repo.GetProductByName(name)
	if err != nil {
		return nil, err
	}
	product.Price = util.ConvertCurrency(product.Price, "INR", currency)
	return product, nil
}

func (s *service) GetProducts() ([]*Product, error) {
	return s.repo.GetProducts()
}

func (s *service) BulkUpdateProducts(products []*Product) (bool, []uint, error) {
	var failedIDs []uint
	for _, product := range products {
		if err := s.repo.UpdateProduct(product); err != nil {
			failedIDs = append(failedIDs, product.ID)
		}
	}
	success := len(failedIDs) == 0
	return success, failedIDs, nil
}
