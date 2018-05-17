package core

import (
	"fmt"
	"log"
)

func DebugLog(msg string, args ...interface{}) {
	if Debug {
		Log(msg, args...)
	}
}

func Log(msg string, args ...interface{}) {
	log.Printf("%s\n", fmt.Sprintf(msg, args...))
}
