package sql

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain01"
)

type (
	Sql struct{}
)

func New() Sql {
	return Sql{}
}

func (s Sql) Save(entity domain01.Domain01) string {
	log.Println("this is just a poc, func: pkg/domain/domain01/repository/sql/sql.go#Save")
	return entity.ID
}

func (s Sql) DeleteByID(id string) {
	log.Println("this is just a poc, func: pkg/domain/domain01/repository/sql/sql.go#DeleteByID")
}
