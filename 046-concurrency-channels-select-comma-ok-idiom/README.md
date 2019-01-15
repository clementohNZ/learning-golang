# Comma Ok Idiom
This pattern, often used for maps to see if a key is available, can also be used to check if a channel is okay 
(i.e. not closed).

[Usage in Maps - Effective Go](https://golang.org/doc/effective_go.html#maps)

## Example 1 - Maps
The below example shows how we would check if the `tz` property exists on the map `timeZone`

```go 
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

The next

## Example 2 - Channels
Here we're looking to see if the channel is still open

```go
case i, ok := <-q:
	// check if channel is okay. If it's not okay then exit
	if !ok {
		fmt.Println("from comma ok", i, ok)
		return
	} else {
		fmt.Println("from comma ok", i)
	}
}
```