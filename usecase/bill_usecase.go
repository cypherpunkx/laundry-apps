package usecase

import (
	"fmt"
	"strings"
	"time"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/repository"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
)

type BillUseCase interface {
	RegisterNewBill(payload model.Bill) error
	FindBillById(id string) (dto.BillResponseDto, error)
}

type billUseCase struct {
	repo repository.BillRepository
	emplUseCase EmployeeUseCase
	custUseCase CustomerUseCase
	prodUseCase ProductUseCase
}

func (b *billUseCase) RegisterNewBill(payload model.Bill) error {
	// get customer
	customer, err := b.custUseCase.FindCustomerById(payload.CustomerId)
	if err != nil {
		return fmt.Errorf("Customer with id %s is not found", payload.CustomerId)
	}

	// get employee
	employee, err := b.emplUseCase.FindEmployeeById(payload.EmployeeId)
	if err != nil {
		return fmt.Errorf("Employee with id %s is not found", payload.EmployeeId)
	}

	newBillDetail := make([]model.BillDetail,0,len(payload.BillDetails))
	for _, billDetail := range payload.BillDetails {
		// getProduct 
		product, err := b.prodUseCase.FindProductById(billDetail.ProductId)
		if err != nil {
			return fmt.Errorf("Product with id %s is not found", billDetail.Id)
		}
		// set selesai laundry base on product yang diambil
		// reguler == 3 hari 
		// expres == 1 hari
		// ToLower => untuk mengubah string menjadi huruf kecil semua
		if strings.ToLower(product.Name) == "reguler" {
			billDetail.FinishDate = time.Now().Add(3 * 24 * time.Hour)
		}else {
			billDetail.FinishDate = time.Now().Add(24 * time.Hour)
		}
		billDetail.Id = common.GenerateUUID()
		billDetail.BillId = payload.Id
		billDetail.ProductId = product.Id
		billDetail.ProductPrice = product.Price
		newBillDetail = append(newBillDetail, billDetail)
	}

	payload.BillDate = time.Now()
	payload.EntryDate= time.Now()
	payload.CustomerId = customer.Id
	payload.EmployeeId = employee.Id
	payload.BillDetails = newBillDetail
	err = b.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to register new bill : %v",err)
	}
	return nil
}

func (b *billUseCase) FindBillById(id string) (dto.BillResponseDto, error) {
	// bill 
	bill, err := b.repo.Get(id)
	if err != nil {
		return dto.BillResponseDto{},fmt.Errorf("Error Get Bill : %s",err.Error())
	}
	// Employee
	employee, err := b.emplUseCase.FindEmployeeById(bill.EmployeeId)
	if err != nil {
		return dto.BillResponseDto{},fmt.Errorf("Error Get Employee : %s",err.Error())
	}
	// Customer
	customer, err := b.custUseCase.FindCustomerById(bill.CustomerId)
	if err != nil {
		return dto.BillResponseDto{},fmt.Errorf("Error Get Customer : %s",err.Error())
	}
	// Get BillDetals
	var billDetailsResponse []dto.BillDetailReponseDto
	// total payment
	var total int
	billDetails , err := b.repo.GetBillDetailByBill(bill.Id)
	if err != nil {
		return dto.BillResponseDto{},fmt.Errorf("Error Get BillDetails : %s",err.Error())
	}
	for _, billDetail := range billDetails {
		var billDetailResponse dto.BillDetailReponseDto
		// Get Product 
		product ,err := b.prodUseCase.FindProductById(billDetail.ProductId)

		if err != nil {
			return dto.BillResponseDto{},fmt.Errorf("Error Get Produt : %s",err.Error())
		}
		billDetailResponse.Product = product
		billDetailResponse.ProductPrice = billDetail.ProductPrice
		total += billDetail.ProductPrice
		billDetailResponse.Qty = billDetail.Qty
		billDetailResponse.FinishDate = billDetail.FinishDate
		billDetailsResponse = append(billDetailsResponse, billDetailResponse)
	}
	// response untuk Billnya
	var billResponse dto.BillResponseDto
	billResponse.BillDate = bill.BillDate
	billResponse.EntryDate = bill.EntryDate
	billResponse.Employee = employee
	billResponse.Customer = customer
	billResponse.BillDetails = billDetailsResponse
	billResponse.TotalBill = total
	return billResponse,nil
}

func NewBillUseCase(
	repository repository.BillRepository,
	employeeUseCase EmployeeUseCase,
	customerUseCase CustomerUseCase,
	productUseCase ProductUseCase,
)BillUseCase{
	return &billUseCase{
		repo: repository,
		emplUseCase: employeeUseCase,
		custUseCase: customerUseCase,
		prodUseCase: productUseCase,
	}
}