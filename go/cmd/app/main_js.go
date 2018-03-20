// +build js
// +build !windows
// +build !linux
// +build !freebsd
// +build !darwin
// +build !android
// +build !ios

package main

import "log"

func init() {
	log.SetFlags(0)
}
