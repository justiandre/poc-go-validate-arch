package service

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain02"
)

//go:generate mockery --all --disable-version-string --exported=true
type (
	repository interface {
		Save(entity domain02.Domain02) string
	}

	Service struct {
		repository repository
	}
)

func New(repository repository) Service {
	return Service{repository: repository}
}

func (s Service) Save(entity domain02.Domain02) string {
	log.Println("this is just a poc, func: pkg/domain/domain02/service/service.go#Save")
	return s.repository.Save(entity)
}
