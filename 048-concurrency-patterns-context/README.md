# Context
[Go Blog - Go Concurrency Patterns: Context](https://blog.golang.org/context)

[Context Package](https://golang.org/pkg/context/)

In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional 
goroutines to access backends such as databases and RPC services. The set of goroutines working on a request 
typically needs access to request-specific values such as the identity of the end user, authorization tokens, 
and the request's deadline. When a request is canceled or times out, all the goroutines working on that request 
should exit quickly so the system can reclaim any resources they are using.

At Google, we developed a context package that makes it easy to pass request-scoped values, cancelation signals, 
and deadlines across API boundaries to all the goroutines involved in handling a request. The package is publicly 
available as context. This article describes how to use the package and provides a complete working example.

Usages: 
1. Stop all launched go routines when you cancel the process that launched them. You don't
want to leak goroutines that still consume resources.
2. Passing around information related to a web request

## Helpful resources
[Understanding the context package](http://p.agnihotry.com/post/understanding_the_context_package_in_golang/)

[Finally some context: Request scoped state in Go 1.7](https://medium.com/@matryer/context-has-arrived-per-request-state-in-go-1-7-4d095be83bd8)
