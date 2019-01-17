# Errors with information
The go library comes with the `errorString` type that allows you to print a red line with some text to indicate an
error. While this may work for a basic application, it's not good for production level logging because you need things
like log levels, timestamps and more.

Check out all the recommended logging libraries from awesome go [here](https://github.com/avelino/awesome-go#logging)