# safe
The quick example demonstrates the idea of how to get mileage out of Safe:
```go
import "safe"

func main() {
    defer safe.CatchExit()
    defer doLongFinalCleanupOperations()
    
    doWork()
    
    if something == wrong {
        safe.Exit(1)
    }
}
```
