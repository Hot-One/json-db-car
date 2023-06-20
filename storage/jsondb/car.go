package jsondb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"

	"app/models"
)

type CarRepo struct {
	fileName string
	file     *os.File
}

func NewCarRepo(fileName string, file *os.File) *CarRepo {
	return &CarRepo{
		fileName: fileName,
		file:     file,
	}
}

func (c *CarRepo) Create(req *models.CarCreate) (*models.Car, error) {
	cars, err := c.read()
	if err != nil {
		return nil, err
	}

	var (
		id  = uuid.New().String()
		car = models.Car{
			Id:    id,
			Model: req.Model,
			Speed: req.Speed,
		}
	)
	cars[id] = car

	err = c.write(cars)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (c *CarRepo) GetById(req *models.CarPrimaryKey) (*models.Car, error) {
	cars, err := c.read()
	if err != nil {
		return nil, err
	}

	if _, have := cars[req.Id]; !have {
		return nil, errors.New("Car not found")
	}

	car := cars[req.Id]
	return &car, nil
}

func (c *CarRepo) GetList(req *models.CarGetListRequest) (*models.CarGetListResponse, error) {
	var resp = &models.CarGetListResponse{}
	resp.Cars = []*models.Car{}
	CarMap, err := c.read()
	if err != nil {
		return nil, err
	}

	resp.Count = len(CarMap)
	for _, val := range CarMap {
		Cars := val
		resp.Cars = append(resp.Cars, &Cars)
	}

	return resp, nil
}

func (c *CarRepo) Update(req *models.CarUpdate) (*models.Car, error) {
	cars, err := c.read()
	if err != nil {
		return nil, err
	}
	cars[req.Id] = models.Car{
		Id:    req.Id,
		Model: req.Model,
		Speed: req.Speed,
	}
	err = c.write(cars)
	if err != nil {
		return nil, err
	}
	car := cars[req.Id]
	return &car, nil
}

func (c *CarRepo) Delete(req *models.CarPrimaryKey) error {
	cars, err := c.read()
	if err != nil {
		return err
	}
	delete(cars, req.Id)

	err = c.write(cars)
	if err != nil {
		return nil
	}
	return nil
}

func (c *CarRepo) read() (map[string]models.Car, error) {
	var (
		cars   []*models.Car
		carMap = make(map[string]models.Car)
	)

	data, err := ioutil.ReadFile(c.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &cars)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, car := range cars {
		carMap[car.Id] = *car
	}

	return carMap, nil
}

func (c *CarRepo) write(carMap map[string]models.Car) error {

	var cars []models.Car

	for _, val := range carMap {
		cars = append(cars, val)
	}

	body, err := json.MarshalIndent(cars, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
