package voerr

import "log"

func Catch() {
	if err := recover(); err != nil {
		log.Println("Catch Error>>:", err)
		return
	}
}
