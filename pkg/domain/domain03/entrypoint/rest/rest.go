package rest

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain02"
	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain03"
)

//go:generate mockery --all --disable-version-string --exported=true
type (
	service interface {
		Save(entity domain03.Domain03, entityDomain02 domain02.Domain02) string
	}

	Rest struct {
		service service
	}
)

func New(service service) Rest {
	return Rest{service: service}
}

func (r Rest) Save(entity domain03.Domain03, entityDomain02 domain02.Domain02) string {
	log.Println("this is just a poc, func: pkg/domain/domain03/entrypoint/rest/rest.go#Save")
	return r.service.Save(entity, entityDomain02)
}
