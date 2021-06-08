API to run code with Docker container

Should the platform be specified as part of the API request? Maybe not! We can maybe bring this up later

What about environment? Of course

Environment - example - Golang v1.16.3
Platform - example - Docker Engine v20.10.6

Something like that. I think it's important to use versions for Platforms too. Previously I didn't do that. It's important because having no version would mean that I support all versions. But I'm gonna be using a particular Platform API which will be compatible only with certain Platform versions, so it only makes sense to start using versions at least in docs or while talking. Maybe in configuration, the user input, in this case it's an admin configuring the server, they can give simply `docker` as input for `platform`

I'm going to start by using some external golang packages for HTTP, or I could just use the standard library, but let's start with less work in trivial things! I was thinking gin or something like that, but I'm going to try out -

https://duckduckgo.com/?t=ffab&q=fast+http&ia=web

https://github.com/valyala/fasthttp

---

Steps
- Run a simple API server first

---

Cool! I was able to run the API server :)

```bash
$ go run main.go
main.go:3:8: no required module provides package github.com/valyala/fasthttp: go.mod file not found in current directory or any parent directory; see 'go help modules'

$ go mod tidy -v
go: go.mod file not found in current directory or any parent directory; see 'go help modules'

$ go mod init github.com/karuppiah7890/run-it/api
go: creating new go.mod: module github.com/karuppiah7890/run-it/api
go: to add module requirements and sums:
	go mod tidy

$ go mod tidy -v
go: finding module for package github.com/valyala/fasthttp
go: found github.com/valyala/fasthttp in github.com/valyala/fasthttp v1.26.0

$ go run main.go
# command-line-arguments
./main.go:6:2: undefined: fmt

$ go run main.go
```

```bash
$ curl localhost:8080
Hi there! RequestURI is "/"
```

:D


