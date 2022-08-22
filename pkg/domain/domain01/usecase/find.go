package usecase

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01"
)

//go:generate mockery --all --disable-version-string --exported=true
type (
	serviceFind interface {
		FindByID(id string) domain01.Domain01
	}

	FindUC struct {
		service serviceFind
	}
)

func NewFindUC(service serviceFind) FindUC {
	return FindUC{service: service}
}

func (uc FindUC) FindByID(id string) domain01.Domain01 {
	log.Println("this is just a poc, func: pkg/domain/domain01/usecase/find.go#FindByID")
	return uc.service.FindByID(id)
}
