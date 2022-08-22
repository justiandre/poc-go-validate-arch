package generic03

import (
	"log"

	"github.com/justiandre/poc-go-validate-arch/pkg/infra/generic02"
)

func GetMsgParse(arg string) string {
	argParse := generic02.Parse(arg)
	log.Println("this is just a poc, func: pkg/infra/generic03/generic03.go#GetMsgParse")
	return argParse
}
