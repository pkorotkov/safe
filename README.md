# safe
The quick example demonstrates the idea of how to get mileage out of Safe:
```go
import "safe"

func main() {
    defer safe.CatchExit()
    defer doLongFinalCleanupOperations()

    go func() {
        err := doWork_1()
        if err != nil {
            safe.Exit(1)
        }
    }()

    go func() {
        err := doWork_2()
        if err != nil {
            safe.Exit(2)
        }
    }()

    WaitForExit()
}
```
