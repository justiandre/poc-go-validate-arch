package sql

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/domain/domain02"
)

type (
	sql struct{}
)

func New() sql {
	return sql{}
}

func (s sql) Save(entity domain02.Domain02) string {
	log.Println("this is just a poc, func: pkg/domain/domain02/repository/sql/sql.go#Save")
	return entity.ID
}
