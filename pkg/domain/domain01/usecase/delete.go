package usecase

import (
	"log"
)

//go:generate mockery --all --disable-version-string --exported=true
type (
	serviceDelete interface {
		DeleteByID(id string)
	}

	DeleteUC struct {
		service serviceDelete
	}
)

func NewDeleteUC(service serviceDelete) DeleteUC {
	return DeleteUC{service: service}
}

func (uc DeleteUC) DeleteByID(id string) {
	log.Println("this is just a poc, func: pkg/domain/domain01/usecase/delete.go#DeleteByID")
	uc.service.DeleteByID(id)
}
