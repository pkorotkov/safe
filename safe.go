package safe // import "github.com/pkorotkov/safe"

import "os"

type exitCode int

// CatchExit catches the final application's exit or rethrows an underlying panic.
func CatchExit() {
	if p := recover(); p != nil {
		if ec, ok := p.(exitCode); ok {
			os.Exit(int(ec))
		}
		panic(p)
	}
}

// Exit triggers application's exit.
// Must be called within CatchExit goroutine.
func Exit(ec int) {
	panic(exitCode(ec))
}
