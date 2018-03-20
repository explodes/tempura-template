package game

import (
	"fmt"
	"log"
)

func DebugLog(msg string, args ...interface{}) {
	if debug {
		Log(msg, args...)
	}
}

func Log(msg string, args ...interface{}) {
	log.Printf("%s\n", fmt.Sprintf(msg, args...))
}
