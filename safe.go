package safe // import "github.com/pkorotkov/safe"

import "os"

type exitCode int

func CatchExit() {
    if p := recover(); p != nil {
        if ec, ok := p.(exitCode); ok {
            os.Exit(int(ec))
        }
        panic(p)
    }
}

func Exit(ec int) {
    panic(exitCode(ec))
}
