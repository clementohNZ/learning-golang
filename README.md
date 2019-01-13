Learning Go
-----------

## Installing go
You can download the version of golang for your operating system [here](https://golang.org/dl/). Make 
sure you select thecorrect system type i.e. 32 or 64 bit else you run into issues down the line.

### Environment Setup
Make sure your:
1. `GOPATH` environment variable is pointing to the proper directory where your go projects
live.
2. `GOROOT` is pointed to the place where you installed go. If you are unsure you can use the command
`which go` to get the path


## How to use
Each folder will contain a `main.go` file. You can run this file using the go cli
and running `go run <path to main.go>` or compiling first using `go build <path to main.go>`.
If you are building the binary using `go build` you will need to run the generated executable
manually. It should be located in the same folder as your `main.go` file, but if it's not there
check your `GOROOT/bin` folder. 