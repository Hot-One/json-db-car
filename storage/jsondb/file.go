package jsondb

import (
	"os"

	"app/config"
	"app/storage"
)

type StoreJSON struct {
	user *UserRepo
	car  *CarRepo
}

func NewConnectionJSON(cfg *config.Config) (storage.StorageI, error) {

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	carFile, err := os.Open(cfg.Path + cfg.CarFileName)
	if err != nil {
		return nil, err
	}

	return &StoreJSON{
		user: NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
		car:  NewCarRepo(cfg.Path+cfg.CarFileName, carFile),
	}, nil
}

func (u *StoreJSON) User() storage.UserRepoI {
	return u.user
}

func (c *StoreJSON) Car() storage.CarRepoI {
	return c.car
}
