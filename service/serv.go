package service

import (
	"github.com/antoha2/buylist/repository"
)

type UnitService interface {
	Create(unit *ServBuyUnit) error
}

type servImpl3 struct {
	rep repository.Repository
	UnitService
}

func NewService(rep repository.Repository) *servImpl3 {
	return &servImpl3{
		rep: rep,
	}
}

type ServBuyUnit struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Text   string `json:"text"`
	IsDone bool   `json:"isDone"`
}
