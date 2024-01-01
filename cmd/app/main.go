package main

import (
	"log"

	"learnbuffalo/actions"
)

// main is the starting point for your Buffalo application.
// You can feel free and add to this `main` method, change
// what it does, etc...
// All we ask is that, at some point, you make sure to
// call `app.Serve()`, unless you don't want to start your
// application that is. :)
func main() {
	app := actions.App()

	// Example of mangaging SSL certs internally w/ Buffalo.
	// serv := servers.Simple{Server: &http.Server{}}
	// s := servers.WrapTLS(serv.Server, "./server.crt", "./server.key")
	// if err := app.Serve(s); err != nil {
	// 	log.Fatal(err)
	// }

	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}

/*
# Notes about `main.go`

## SSL Support

We recommend placing your application behind a proxy, such as
Apache or Nginx and letting them do the SSL heavy lifting
for you. https://gobuffalo.io/en/docs/proxy

## Buffalo Build

When `buffalo build` is run to compile your binary, this `main`
function will be at the heart of that binary. It is expected
that your `main` function will start your application using
the `app.Serve()` method.

*/
