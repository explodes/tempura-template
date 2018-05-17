package core

const (
	// -ldflags "-X main.Debug=false" or "-X mobile.Debug=false"
	// does not appear to be supported with gopherjs and gomobile.
	//
	// The makefile handles this with sed.
	//
	// The default state should be "true" for source control, but
	// all builds should specify their debug mode.
	Debug = true
)
