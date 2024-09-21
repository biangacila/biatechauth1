package adapters

import (
	"github.com/biangacila/biatechauth1/domain/aggregates"
	"github.com/biangacila/biatechauth1/domain/repositories"
	"sync"
)

type MemoryProductRepository struct {
	products map[string]aggregates.ProductAggregate
	sync.Mutex
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[string]aggregates.ProductAggregate),
	}
}

func (m *MemoryProductRepository) GetAll() ([]aggregates.ProductAggregate, error) {
	var products []aggregates.ProductAggregate
	for _, product := range m.products {
		products = append(products, product)
	}
	return products, nil
}

func (m *MemoryProductRepository) GetByID(id string) (aggregates.ProductAggregate, error) {
	if product, ok := m.products[id]; ok {
		return product, nil
	}
	return aggregates.ProductAggregate{}, repositories.ErrProductNotFound
}

func (m *MemoryProductRepository) Add(product aggregates.ProductAggregate) (aggregates.ProductAggregate, error) {
	m.Lock()
	defer m.Unlock()

	if m.products == nil {
		m.products = make(map[string]aggregates.ProductAggregate)
	}
	m.products[product.AggregateID()] = product
	return product, nil
}

func (m *MemoryProductRepository) Update(product aggregates.ProductAggregate) (aggregates.ProductAggregate, error) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[product.AggregateID()]; !ok {
		return aggregates.ProductAggregate{}, repositories.ErrProductNotFound
	}
	m.products[product.AggregateID()] = product
	return product, nil
}

func (m *MemoryProductRepository) Delete(id string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[id]; !ok {
		return repositories.ErrProductNotFound
	}
	delete(m.products, id)
	return nil
}
