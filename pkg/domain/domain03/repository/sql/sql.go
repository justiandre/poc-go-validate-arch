package sql

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain03"
)

type (
	Sql struct{}
)

func New() Sql {
	return Sql{}
}

func (s Sql) Save(entity domain03.Domain03) string {
	log.Println("this is just a poc, func: pkg/domain/domain03/repository/sql/sql.go#Save")
	return entity.ID
}
