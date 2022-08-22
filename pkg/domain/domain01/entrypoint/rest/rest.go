package rest

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01"
)

//go:generate mockery --all --disable-version-string --exported=true
type (
	service interface {
		FindByID(id string) domain01.Domain01
		Save(entity domain01.Domain01) string
		DeleteByID(id string)
	}

	Rest struct {
		service service
	}
)

func New(service service) Rest {
	return Rest{service: service}
}

func (r Rest) FindByID(id string) domain01.Domain01 {
	log.Println("this is just a poc, func: pkg/domain/domain01/entrypoint/rest/rest.go#FindByID")
	return r.service.FindByID(id)
}

func (r Rest) Save(entity domain01.Domain01) string {
	log.Println("this is just a poc, func: pkg/domain/domain01/entrypoint/rest/rest.go#Save")
	return r.service.Save(entity)
}

func (r Rest) DeleteByID(id string) {
	log.Println("this is just a poc, func: pkg/domain/domain01/entrypoint/rest/rest.go#DeleteByID")
	r.service.DeleteByID(id)
}
