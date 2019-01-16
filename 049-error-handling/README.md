## Basic Custom Errors
If your struct implements the method `Error()` seen in the interface below, it will automatically be
recognized as an error type.

Need to get used to the auto-implementation of interfaces in Go.

```go
type error interface {
    Error() string
}
```

The error built-in interface type is the conventional interface for representing an error condition, 
with the nil value representing no error.

##Links
[Why doesn't go have exceptions?](https://golang.org/doc/faq#exceptions)

[Go Blog - Error Handling](https://blog.golang.org/error-handling-and-gor)

[Go Blog - Defer, Panic, and Recover](https://blog.golang.org/defer-panic-and-recover)

[Go By Example - Errors](https://gobyexample.com/errors)

[Go By Example - Defer](https://gobyexample.com/defer)

[Go By Example - Panic](https://gobyexample.com/panic)

[Golang Bot - Panic and Recover](https://golangbot.com/panic-and-recover/)

[Golang Bot - Custom Errors](https://golangbot.com/custom-errors/)