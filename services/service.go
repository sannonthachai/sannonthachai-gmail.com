package services

import (
	"github.com/sannonthachai/test/domains"
	"github.com/sannonthachai/test/models"
	"github.com/sannonthachai/test/repositories"
)

type Service struct {
	Serv *repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{
		Serv: repo,
	}
}

func (serv *Service) GetItem() []domains.TableSchema {
	res := serv.Serv.FindAll()
	return res
}

func (serv *Service) SaveItem(request *models.Model) {
	serv.Serv.Create(request.Code, request.Price)
}

func (serv *Service) DeleteItem(code string) {
	serv.Serv.DeleteBy(code)
}
