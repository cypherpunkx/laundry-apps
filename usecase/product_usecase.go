package usecase

import (
	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/repository"	
)

type ProductUseCase interface {
	RegisterNewProduct(payload model.Product) error
	FindAllProduct(requestPaging dto.PaginationParam) ([]model.Product,dto.Paging,error) 
	FindProductById(id string) (model.Product,error)
}

type productUseCase struct {
	repo repository.ProductRepository
}

func (p *productUseCase) RegisterNewProduct(payload model.Product) error {
	return p.repo.Create(payload)
}

func (p *productUseCase) FindAllProduct(requestPaging dto.PaginationParam) ([]model.Product,dto.Paging,error)  {
	return p.repo.Paging(requestPaging)
}

func (p *productUseCase) FindProductById(id string) (model.Product,error) {
	return p.repo.Get(id)
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		repo: repository,
	}
}