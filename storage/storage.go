package storage

import "app/models"

type StorageI interface {
	User() UserRepoI
	Car() CarRepoI
}

type UserRepoI interface {
	Create(req *models.CreateUser) (*models.User, error)
	GetById(req *models.UserPrimaryKey) (*models.User, error)
	GetList(req *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(req *models.UserUpdate) (*models.User, error)
	Delete(req *models.UserPrimaryKey) error
}

type CarRepoI interface {
	Create(req *models.CarCreate) (*models.Car, error)
	GetById(req *models.CarPrimaryKey) (*models.Car, error)
	GetList(req *models.CarGetListRequest) (*models.CarGetListResponse, error)
	Update(req *models.CarUpdate) (*models.Car, error)
	Delete(req *models.CarPrimaryKey) error
}
