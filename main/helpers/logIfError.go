package helpers

import "log"

func LogIfError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
