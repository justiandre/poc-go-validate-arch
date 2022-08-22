package generic02

import "log"

func Parse(arg string) string {
	msg := "this is just a poc, func: pkg/infra/generic02/generic03.go#Parse, parse:" + arg
	log.Println(msg)
	return msg
}
