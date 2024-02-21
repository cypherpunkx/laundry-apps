package manager

import "enigmacamp.com/enigma-laundry-apps/usecase"

type UseCaseManager interface {
	ProductUseCase() usecase.ProductUseCase
	CustomerUseCase() usecase.CustomerUseCase
	EmployeeUseCase() usecase.EmployeeUseCase
	BillUseCase() usecase.BillUseCase
	UserUseCase() usecase.UserUseCase
	AuthUseCase() usecase.AuthUseCase
	UserPictureUseCase() usecase.UserPictureUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepo())
}

func (u *useCaseManager) EmployeeUseCase() usecase.EmployeeUseCase {
	return usecase.NewEmployeeUseCase(u.repoManager.EmployeeRepo())
}

func (u *useCaseManager) BillUseCase() usecase.BillUseCase {
	return usecase.NewBillUseCase(u.repoManager.BillRepo(),u.EmployeeUseCase(),u.CustomerUseCase(),u.ProductUseCase())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.UserUseCase())
}

func ( u *useCaseManager) UserPictureUseCase() usecase.UserPictureUseCase {
	return usecase.NewUserPictureUseCase(u.repoManager.UserPicture(),u.repoManager.FileRepository(),u.UserUseCase())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}