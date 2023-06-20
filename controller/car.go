package controller

import (
	"log"

	"app/models"
)

func (c *Controller) CarCreate(req *models.CarCreate) (*models.Car, error) {

	log.Printf("Car create request: %+v\n", req)

	resp, err := c.Strg.Car().Create(req)
	if err != nil {
		log.Printf("Error while car create: %+v\n", err)
		return nil, err
	}
	return resp, nil
}

func (c *Controller) CarGetById(req *models.CarPrimaryKey) (*models.Car, error) {
	resp, err := c.Strg.Car().GetById(req)
	if err != nil {
		log.Printf("Error while car GetById: %+v", err)
		return nil, err
	}
	return resp, nil
}

func (c *Controller) CarGetList(req *models.CarGetListRequest) (*models.CarGetListResponse, error) {
	resp, err := c.Strg.Car().GetList(req)
	if err != nil {
		log.Printf("Error while car GetList: %+v", err)
		return nil, err
	}
	return resp, nil
}

func (c *Controller) CarUpdate(req *models.CarUpdate) (*models.Car, error) {
	resp, err := c.Strg.Car().Update(req)
	if err != nil {
		log.Printf("Error while car Update: %+v", err)
		return nil, err
	}

	return resp, nil
}

func (c Controller) CarDelete(req *models.CarPrimaryKey) error {
	err := c.Strg.Car().Delete(req)
	if err != nil {
		log.Printf("Error while car Delete: %+v", err)
		return err
	}
	return nil
}
