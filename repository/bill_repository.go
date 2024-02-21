package repository

import (
	"database/sql"
	"fmt"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/utils/constant"
)

type BillRepository interface {
	Create(payload model.Bill) error 
	Get(id string) (model.Bill, error)
	List(requestPagin dto.PaginationParam) ([]dto.BillResponseDto, dto.Paging,error)
	GetBillDetailByBill(id string) ([]model.BillDetail,error)
}

type billRepository struct {
	db *sql.DB
}

func (b *billRepository) Create(payload model.Bill) error  {
	tx, err := b.db.Begin()
	if err != nil {		
		return err
	}
	// insert bill
	_,err = tx.Exec(
		constant.BILL_CREATE,
		payload.Id,
		payload.BillDate,
		payload.EntryDate,		
		payload.EmployeeId,
		payload.CustomerId,
	)

	if err != nil {
		return err
	}

	for _,item := range payload.BillDetails {
		_, err = tx.Exec(constant.BIll_DETAIL_CREATE,item.Id,item.BillId,item.ProductId,item.ProductPrice,item.Qty,item.FinishDate) 
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (b *billRepository) Get(id string) (model.Bill, error) {
	var bill model.Bill
	err := b.db.QueryRow(constant.BILL_GET,id).Scan(
		&bill.Id,
		&bill.BillDate,
		&bill.EntryDate,
		&bill.EmployeeId,
		&bill.CustomerId,
	)
	if err != nil {
		return model.Bill{},fmt.Errorf("Error Get Bill : %s ", err.Error())
	}
	return bill,nil
}

func ( b *billRepository) List(requestPagin dto.PaginationParam) ([]dto.BillResponseDto, dto.Paging,error){
	return nil,dto.Paging{},nil
}

func (b *billRepository) GetBillDetailByBill(id string) ([]model.BillDetail,error) {
	var billDetails []model.BillDetail
	rows,err := b.db.Query(constant.BIll_DETAIL_GET,id)
	if err != nil {
		return nil,err
	}
	for rows.Next() {
		var billDetail model.BillDetail
		err := rows.Scan(
			&billDetail.Id,
			&billDetail.BillId,
			&billDetail.ProductId,
			&billDetail.ProductPrice,
			&billDetail.Qty,
			&billDetail.FinishDate,
		)
		if err != nil {
			return nil,err
		}
		billDetails = append(billDetails, billDetail)
	}
	if err != nil {
		return nil,err
	}
	return billDetails,nil
}

func NewBillRepository(db *sql.DB) BillRepository {
	return &billRepository{
		db : db,
	}
}