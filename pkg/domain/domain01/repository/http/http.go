package http

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01"
)

type (
	Http struct{}
)

func New() Http {
	return Http{}
}

func (h Http) FindByID(id string) domain01.Domain01 {
	log.Println("this is just a poc, func: pkg/domain/domain01/repository/http/http.go#FindByID")
	return domain01.Domain01{ID: id, Value01: "a", Value02: "b"}
}
