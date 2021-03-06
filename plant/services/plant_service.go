package services

import (
	"plant/entities"
	"plant/repositories"
)

type PlantService struct {}

func GetService() *PlantService {
	return &PlantService{}
}

func (b *PlantService) FindAll() []entities.Plant{
	return repositories.FindAll()
}
