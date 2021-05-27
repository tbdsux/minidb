package minidb

import "log"

// log the message if err != nil
func logError(err error, message interface{}) {
	if err != nil {
		log.Fatalln(message)
	}
}
