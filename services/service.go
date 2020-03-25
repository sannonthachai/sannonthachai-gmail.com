package services

import (
	"github.com/sannonthachai/test/domains"
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
