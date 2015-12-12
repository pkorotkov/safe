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

// Exit is goroutine-safe application exit trigger.
func Exit(ec int) {
    exit <- exitCode(ec)
}

// WaitExit blocks and waits for an exit signal and triggers the final application's exit.
// This function must be invoked in the same goroutine as CatchExit (typically, at the main's end).
func WaitForExit() {
    select {
    case ec := <- exit:
        panic(ec)
    }
}
