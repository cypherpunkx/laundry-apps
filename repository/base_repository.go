package repository

import "enigmacamp.com/enigma-laundry-apps/model/dto"

type BaseRepository[T any] interface {
	Create(payload T) error
	List() ([]T, error)
	Get(id string) (T, error)
	Update(payload T) error
	Delete(id string) error
}

type BaseRepositoryPaging[T any] interface {
	Paging(requestPaging dto.PaginationParam, query ...string) ([]T, dto.Paging, error)
}
