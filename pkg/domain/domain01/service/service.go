package service

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01"
)

//go:generate mockery --all --disable-version-string --exported=true
type (
	repository interface {
		FindByID(id string) domain01.Domain01
		Save(entity domain01.Domain01) string
		DeleteByID(id string)
	}

	Service struct {
		repository repository
	}
)

func New(repository repository) Service {
	return Service{repository: repository}
}

func (s Service) FindByID(id string) domain01.Domain01 {
	log.Println("this is just a poc, func: pkg/domain/domain01/service/service.go#FindByID")
	return s.repository.FindByID(id)
}

func (s Service) Save(entity domain01.Domain01) string {
	log.Println("this is just a poc, func: pkg/domain/domain01/service/service.go#Save")
	return s.repository.Save(entity)
}

func (s Service) DeleteByID(id string) {
	log.Println("this is just a poc, func: pkg/domain/domain01/service/service.go#DeleteByID")
	s.repository.DeleteByID(id)
}
