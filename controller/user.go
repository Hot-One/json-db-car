package controller

import (
	"errors"
	"log"

	"app/models"
)

func (c *Controller) UserCreate(req *models.CreateUser) (*models.User, error) {

	log.Printf("User create req: %+v\n", req)

	resp, err := c.Strg.User().Create(req)
	if err != nil {
		log.Printf("Error while user Create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) UserGetById(req *models.UserPrimaryKey) (*models.User, error) {
	resp, err := c.Strg.User().GetById(req)
	if err != nil {
		log.Printf("Error while user GetById: %+v\n", err)
		return nil, err
	}
	return resp, nil
}

func (c *Controller) UserGetList(req *models.UserGetListRequest) (*models.UserGetListResponse, error) {
	resp, err := c.Strg.User().GetList(req)
	if err != nil {
		log.Printf("Error while user GetList: %+v\n", err)
	}
	return resp, nil
}

func (c *Controller) UserUpdate(req *models.UserUpdate) (*models.User, error) {
	resp, err := c.Strg.User().Update(req)
	if err != nil {
		log.Printf("Error while user Update: %+v\n", err)
		return nil, err
	}
	return resp, nil
}

func (c *Controller) UserDelete(req *models.UserPrimaryKey) *models.User {
	err := c.Strg.User().Delete(req)
	if err != nil {
		log.Printf("Error while user Delete: %+v\n", err)
	}
	return nil
}
