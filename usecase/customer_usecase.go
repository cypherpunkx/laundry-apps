package usecase

import (
	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/repository"
)

type CustomerUseCase interface {
	RegisterNewCustomer(payload model.Customer) error
	FindAllCustomer(requestPaging dto.PaginationParam) ([]model.Customer,dto.Paging,error) 
	FindCustomerById(id string) (model.Customer,error)
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (c *customerUseCase) RegisterNewCustomer(payload model.Customer) error {
	return c.repo.Create(payload)
}

func (c *customerUseCase) FindAllCustomer(requestPaging dto.PaginationParam) ([]model.Customer,dto.Paging,error)  {
	return c.repo.Paging(requestPaging)
}

func (c *customerUseCase) FindCustomerById(id string) (model.Customer,error) {
	return c.repo.Get(id)
}


func NewCustomerUseCase(repository repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		repo: repository,
	}
}