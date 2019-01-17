# Basic Logging
This is logging the go standard library comes with. There are other logging libraries out there that 
will be more feature-complete.

## Key Points
- When `fatal`, deferred functions usually don't run because the program exits.
- When `panic` is involved, usually means you can recover from them and **deferred 
functions still run**.

## Panic
### Go Function Definition
`func panic(v interface{})`

The panic built-in function stops normal execution of the current goroutine. When a function F 
calls panic, normal execution of F stops immediately. Any functions whose execution was deferred 
by F are run in the usual way, and then F returns to its caller. To the caller G, the invocation 
of F then behaves like a call to panic, terminating G's execution and running any deferred 
functions. This continues until all functions in the executing goroutine have stopped, in reverse 
order. At that point, the program is terminated and the error condition is reported, including the 
value of the argument to panic. This termination sequence is called panicking and can be controlled 
by the built-in function recover.

### Effective Go
The usual way to report an error to a caller is to return an error as an extra return value. 
The canonical Read method is a well-known instance; it returns a byte count and an error. 
But what if the error is unrecoverable? Sometimes the program simply cannot continue.

For this purpose, there is a built-in function panic that in effect creates a run-time error 
that will stop the program (but see the next section). The function takes a single argument of 
arbitrary type—often a string—to be printed as the program dies. It's also a way to indicate that 
something impossible has happened, such as exiting an infinite loop.

```go
// A toy implementation of cube root using Newton's method.
func CubeRoot(x float64) float64 {
    z := x/3   // Arbitrary initial value
    for i := 0; i < 1e6; i++ {
        prevz := z
        z -= (z*z*z-x) / (3*z*z)
        if veryClose(z, prevz) {
            return z
        }
    }
    // A million iterations has not converged; something is wrong.
    panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}
```

This is only an example but real library functions should avoid panic. If the problem can be masked 
or worked around, it's always better to let things continue to run rather than taking down the whole 
program. One possible counterexample is during initialization: if the library truly cannot set itself 
up, it might be reasonable to panic, so to speak.

```go
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```

### Golang Spec - Handling Panics
Handling panics
Two built-in functions, panic and recover, assist in reporting and handling run-time panics and 
program-defined error conditions.

func panic(interface{})
func recover() interface{}
While executing a function F, an explicit call to panic or a run-time panic terminates the execution of F. 
Any functions deferred by F are then executed as usual. Next, any deferred functions run by F's caller are run, 
and so on up to any deferred by the top-level function in the executing goroutine. At that point, the program is 
terminated and the error condition is reported, including the value of the argument to panic. This termination 
sequence is called panicking.

```go
panic(42)
panic("unreachable")
panic(Error("cannot parse"))
```
The recover function allows a program to manage behavior of a panicking goroutine. Suppose a function 
G defers a function D that calls recover and a panic occurs in a function on the same goroutine in which 
G is executing. When the running of deferred functions reaches D, the return value of D's call to recover 
will be the value passed to the call of panic. If D returns normally, without starting a new panic, the 
panicking sequence stops. In that case, the state of functions called between G and the call to panic is 
discarded, and normal execution resumes. Any functions deferred by G before D are then run and G's execution 
terminates by returning to its caller.

The return value of recover is nil if any of the following conditions holds:

- panic's argument was nil;
- the goroutine is not panicking;
- recover was not called directly by a deferred function.

The protect function in the example below invokes the function argument g and protects callers from run-time panics raised by g.

```go
func protect(g func()) {
	defer func() {
		log.Println("done")  // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}
```

## Recover
### Go Function Definition
`func recover() interface{}`
The recover built-in function allows a program to manage behavior of a panicking goroutine. Executing a call to 
recover inside a deferred function (but not any function called by it) stops the panicking sequence by restoring 
normal execution and retrieves the error value passed to the call of panic. If recover is called outside the 
deferred function it will not stop a panicking sequence. In this case, or when the goroutine is not panicking, or 
if the argument supplied to panic was nil, recover returns nil. Thus the return value from recover reports whether 
the goroutine is panicking.

### Effective Go
When panic is called, including implicitly for run-time errors such as indexing a slice out of bounds or 
failing a type assertion, it immediately stops execution of the current function and begins unwinding the 
stack of the goroutine, running any deferred functions along the way. If that unwinding reaches the top of 
the goroutine's stack, the program dies. However, it is possible to use the built-in function recover to 
regain control of the goroutine and resume normal execution.

A call to recover stops the unwinding and returns the argument passed to panic. Because the only code that 
runs while unwinding is inside deferred functions, recover is only useful inside deferred functions.

One application of recover is to shut down a failing goroutine inside a server without killing the other 
executing goroutines.

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

In this example, if do(work) panics, the result will be logged and the goroutine will exit cleanly without 
disturbing the others. There's no need to do anything else in the deferred closure; calling recover handles 
the condition completely.

Because recover always returns nil unless called directly from a deferred function, deferred code can call 
library routines that themselves use panic and recover without failing. As an example, the deferred function 
in safelyDo might call a logging function before calling recover, and that logging code would run unaffected 
by the panicking state.

With our recovery pattern in place, the do function (and anything it calls) can get out of any bad situation 
cleanly by calling panic. We can use that idea to simplify error handling in complex software. Let's look at 
an idealized version of a regexp package, which reports parsing errors by calling panic with a local error type. 
Here's the definition of Error, an error method, and the Compile function.

```go
// Error is the type of a parse error; it satisfies the error interface.
type Error string
func (e Error) Error() string {
    return string(e)
}

// error is a method of *Regexp that reports parsing errors by
// panicking with an Error.
func (regexp *Regexp) error(err string) {
    panic(Error(err))
}

// Compile returns a parsed representation of the regular expression.
func Compile(str string) (regexp *Regexp, err error) {
    regexp = new(Regexp)
    // doParse will panic if there is a parse error.
    defer func() {
        if e := recover(); e != nil {
            regexp = nil    // Clear return value.
            err = e.(Error) // Will re-panic if not a parse error.
        }
    }()
    return regexp.doParse(str), nil
}
```

If doParse panics, the recovery block will set the return value to nil—deferred functions can modify named 
return values. It will then check, in the assignment to err, that the problem was a parse error by asserting 
that it has the local type Error. If it does not, the type assertion will fail, causing a run-time error that 
continues the stack unwinding as though nothing had interrupted it. This check means that if something unexpected 
happens, such as an index out of bounds, the code will fail even though we are using panic and recover to handle 
parse errors.

With error handling in place, the error method (because it's a method bound to a type, it's fine, even natural, 
for it to have the same name as the builtin error type) makes it easy to report parse errors without worrying 
about unwinding the parse stack by hand:

```go
if pos == 0 {
    re.error("'*' illegal at start of expression")
}
```

Useful though this pattern is, it should be used only within a package. Parse turns its internal panic calls into 
error values; it does not expose panics to its client. That is a good rule to follow.

By the way, this re-panic idiom changes the panic value if an actual error occurs. However, both the original and 
new failures will be presented in the crash report, so the root cause of the problem will still be visible. Thus 
this simple re-panic approach is usually sufficient—it's a crash after all—but if you want to display only the 
original value, you can write a little more code to filter unexpected problems and re-panic with the original 
error. That's left as an exercise for the reader.