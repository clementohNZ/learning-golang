# Documenting Code
Go generates documents automatically based on your comments so formatting is crucial.

## View unpublished documentation locally
The publish–review–fix–republish cycle can be pretty tedious, and it makes for a noisy commit log.

A great way to short-circuit this cycle is to run a godoc server locally so that you can click around the 
documentation for your project in your own browser.

Pick a port to run the server on, say 6060, and start it with the following command:

```sh
godoc -http=":6060"
```

Then you can browse all the installed packages at localhost:6060/pkg. You don’t have to restart the server to 
get the changes when you edit a doc comment, just save the file and refresh the page in the browser.

## Helpful Links
[Godoc - Godoc Tricks](https://godoc.org/github.com/fluhus/godoc-tricks)