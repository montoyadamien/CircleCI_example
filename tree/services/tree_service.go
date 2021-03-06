package services

import (
	"tree/entities"
	"tree/repositories"
)

type PlantService struct {}

func GetService() *PlantService {
	return &PlantService{}
}

func (b *PlantService) FindAll() []entities.Tree{
	return repositories.FindAll()
}
