package repository

import (
	"database/sql"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"enigmacamp.com/enigma-laundry-apps/utils/constant"
)

type ProductRepository interface {
	BaseRepository[model.Product]
	BaseRepositoryPaging[model.Product]
}	

type productRepository struct {
	db *sql.DB
}

func (p *productRepository) Create(payload model.Product) error {
	_,err := p.db.Exec(constant.PRODUCT_INSERT,payload.Id,payload.Name,payload.Price,payload.Uom)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) List()([]model.Product,error){
	return nil,nil
}

func (p *productRepository) Get(id string) (model.Product,error){
	var product model.Product
	row := p.db.QueryRow(constant.PRODUCT_GET,id)
	err := row.Scan(&product.Id,&product.Name, &product.Price, &product.Uom)
	if err != nil {
		return model.Product{},err
	}
	return product,nil
}

func (p *productRepository) Update(payload model.Product) error {
	_, err := p.db.Exec(constant.PRODUCT_UPDATE,payload.Name,payload.Price,payload.Uom,payload.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) Delete(id string ) error {
	_, err := p.db.Exec(constant.PRODUCT_DELETE,id)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) Paging(requestPaging dto.PaginationParam, query ...string) ([]model.Product, dto.Paging, error) {
	var paginationQuery dto.PaginationQuery
	paginationQuery = common.GetPaginationParams(requestPaging)
	querySelect := "SELECT * FROM product LIMIT $1 OFFSET $2"	
	rows, err := p.db.Query(querySelect, paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Uom)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		products = append(products, product)
	}

	// count total rows
	var totalRows int
	row := p.db.QueryRow("SELECT COUNT(*) FROM product")
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	return products, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db:db}
}