package safe

import (
    "testing"
    "time"
)

func TestMain(m *testing.M) {
    defer CatchExit()
    go func() {
        time.Sleep(2 * time.Second)
        Exit(1)
    }()
    go func() {
        time.Sleep(2 * time.Second)
        Exit(2)
    }()
    go func() {
        time.Sleep(1 * time.Second)
        Exit(0)
    }()
    WaitExit()
}
