package minidb

import "log"

func logError(err error, message interface{}) {
	if err != nil {
		log.Fatalln(message)
	}
}
