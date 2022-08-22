package usecase

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01"
)

//go:generate mockery --all --disable-version-string --exported=true
type (
	serviceSave interface {
		Save(entity domain01.Domain01) string
	}

	SaveUC struct {
		service serviceSave
	}
)

func NewSaveUC(service serviceSave) SaveUC {
	return SaveUC{service: service}
}

func (uc SaveUC) Save(entity domain01.Domain01) string {
	log.Println("this is just a poc, func: pkg/domain/domain01/usecase/save.go#Save")
	return uc.service.Save(entity)
}
