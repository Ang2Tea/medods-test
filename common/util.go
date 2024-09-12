package common

import "log"

func PanicIfErr(err error) {
	if err != nil {
		log.Panicln(err.Error())
	}
}
