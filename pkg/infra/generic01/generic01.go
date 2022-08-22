package generic01

import "log"

func GetMsg() string {
	msg := "this is just a poc, func: pkg/infra/generic01/generic01.go#GetMsg"
	log.Println(msg)
	return msg
}
