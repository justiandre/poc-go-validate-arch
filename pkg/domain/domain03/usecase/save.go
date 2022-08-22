package usecase

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain02"
	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain03"
)

//go:generate mockery --all --disable-version-string --exported=true
type (
	serviceSave interface {
		Save(entity domain03.Domain03) string
	}

	serviceDomain02Save interface {
		Save(entity domain02.Domain02) string
	}

	SaveUC struct {
		service             serviceSave
		serviceDomain02Save serviceDomain02Save
	}
)

func NewSaveUC(service serviceSave, serviceDomain02Save serviceDomain02Save) SaveUC {
	return SaveUC{service: service, serviceDomain02Save: serviceDomain02Save}
}

func (uc SaveUC) Save(entity domain03.Domain03, entityDomain02 domain02.Domain02) string {
	log.Println("this is just a poc, func: pkg/domain/domain03/usecase/save.go#Save")
	uc.serviceDomain02Save.Save(entityDomain02)
	return uc.service.Save(entity)
}
