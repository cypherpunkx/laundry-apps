package manager

import "enigmacamp.com/enigma-laundry-apps/repository"

type RepoManager interface {
	BillRepo() repository.BillRepository
	CustomerRepo() repository.CustomerRepository
	EmployeeRepo() repository.EmployeeRepository
	ProductRepo() repository.ProductRepository
	UserRepo() repository.UserRepository
	UserPicture() repository.UserPicture
	FileRepository() repository.FileRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infra.Conn())
}

func (r *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Conn())
}

func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infra.Conn())
}

func (r *repoManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.Conn())
}

func (r *repoManager)  UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func ( r *repoManager) UserPicture() repository.UserPicture {
	return repository.NewUserPicrerRepository(r.infra.Conn())
}

func (r *repoManager) FileRepository() repository.FileRepository {
	return repository.NewFileRepository(r.infra.GetConfig().FileConfig.UserPicturePath)
}

func NewRepoManager(infraParam InfraManager) RepoManager {
	return &repoManager{
		infra: infraParam,
	}
}