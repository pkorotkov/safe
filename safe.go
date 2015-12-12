package safe // import "github.com/pkorotkov/safe"

import "os"

type exitCode int

var exit = make(chan exitCode)

// CatchExit catches the final application's exit or rethrows an underlying panic.
func CatchExit() {
    if p := recover(); p != nil {
        if ec, ok := p.(exitCode); ok {
            os.Exit(int(ec))
        }
        panic(p)
    }
}

// Exit triggers application exit in goroutine-safe manner.
// Must not be called within Wait-Catch goroutine.
func Exit(ec int) {
    exit <- exitCode(ec)
}

// WaitForExit blocks and waits for an exit signal and triggers the final application's exit.
// This function must be invoked in the same goroutine as CatchExit (typically, at the main's end).
func WaitForExit() {
    select {
    case ec := <- exit:
        panic(ec)
    }
}
