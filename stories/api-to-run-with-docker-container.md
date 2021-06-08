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

---

Next steps:
- Using an API request invoke a Docker container. Maybe use golang image. Maybe run a echo command, or just `go version` command

---

I need to keep the platform code separately so that I can easily refactor later, using interfaces etc :) I don't want to refactor too soon, but also not keep everything very tightly coupled already if it can be separate simple functions :)

https://docs.docker.com/engine/api/sdk/

```bash
$ go get -u -v github.com/docker/docker/client
go: downloading github.com/docker/docker v1.13.1
go: downloading github.com/docker/docker v20.10.7+incompatible
go: downloading github.com/docker/go-units v0.4.0
go: downloading github.com/docker/distribution v2.7.1+incompatible
go: downloading github.com/docker/go-connections v0.4.0
go: downloading github.com/gogo/protobuf v1.3.2
go: downloading github.com/containerd/containerd v1.5.2
go: downloading github.com/sirupsen/logrus v1.8.1
go: downloading github.com/opencontainers/image-spec v1.0.1
go: downloading github.com/opencontainers/go-digest v1.0.0
go: downloading google.golang.org/grpc v1.38.0
go: downloading golang.org/x/sys v0.0.0-20210514084401-e8d321eab015
go: downloading github.com/Microsoft/go-winio v0.4.17
go: downloading golang.org/x/net v0.0.0-20210510120150-4163338589ed
go: downloading google.golang.org/genproto v0.0.0-20201110150050-8816d57aaa9a
go: downloading github.com/golang/protobuf v1.4.3
go: downloading github.com/golang/protobuf v1.5.2
go: downloading github.com/Microsoft/go-winio v0.5.0
go: downloading google.golang.org/protobuf v1.25.0
go: downloading golang.org/x/sys v0.0.0-20210603125802-9665404d3644
go: downloading golang.org/x/net v0.0.0-20210525063256-abc453219eb5
go: downloading google.golang.org/protobuf v1.26.0
go: downloading google.golang.org/genproto v0.0.0-20210607140030-00d4fb20b1ae
google.golang.org/protobuf/internal/set
google.golang.org/protobuf/internal/flags
github.com/docker/docker/api/types/events
github.com/docker/docker/api
golang.org/x/sys/internal/unsafeheader
github.com/docker/docker/api/types/image
google.golang.org/protobuf/internal/pragma
github.com/docker/docker/api/types/versions
golang.org/x/sys/unix
google.golang.org/protobuf/internal/detrand
google.golang.org/grpc/codes
github.com/pkg/errors
google.golang.org/protobuf/internal/version
github.com/opencontainers/go-digest
github.com/opencontainers/image-spec/specs-go
google.golang.org/protobuf/internal/errors
github.com/docker/docker/api/types/blkiodev
github.com/docker/docker/api/types/mount
github.com/docker/docker/api/types/strslice
github.com/docker/go-connections/nat
github.com/docker/go-units
github.com/docker/docker/api/types/filters
google.golang.org/protobuf/encoding/protowire
github.com/docker/docker/api/types/time
github.com/gogo/protobuf/proto
google.golang.org/protobuf/reflect/protoreflect
golang.org/x/net/internal/socks
github.com/docker/docker/api/types/container
github.com/docker/distribution/digestset
github.com/opencontainers/image-spec/specs-go/v1
github.com/docker/docker/api/types/network
github.com/docker/docker/api/types/registry
github.com/docker/distribution/registry/api/errcode
golang.org/x/net/proxy
github.com/docker/go-connections/tlsconfig
net/http/httputil
github.com/docker/distribution/reference
github.com/docker/go-connections/sockets
google.golang.org/protobuf/internal/encoding/messageset
google.golang.org/protobuf/internal/strs
google.golang.org/protobuf/internal/descfmt
google.golang.org/protobuf/runtime/protoiface
google.golang.org/protobuf/internal/order
google.golang.org/protobuf/internal/genid
google.golang.org/protobuf/internal/encoding/text
google.golang.org/protobuf/internal/descopts
google.golang.org/protobuf/reflect/protoregistry
github.com/sirupsen/logrus
google.golang.org/protobuf/internal/encoding/defval
google.golang.org/protobuf/proto
github.com/containerd/containerd/log
google.golang.org/protobuf/encoding/prototext
google.golang.org/protobuf/internal/filedesc
google.golang.org/protobuf/internal/encoding/tag
google.golang.org/protobuf/internal/impl
github.com/docker/docker/api/types/swarm/runtime
github.com/docker/docker/api/types/swarm
github.com/docker/docker/api/types
github.com/docker/docker/api/types/volume
google.golang.org/protobuf/internal/filetype
google.golang.org/protobuf/runtime/protoimpl
google.golang.org/protobuf/types/known/anypb
google.golang.org/protobuf/types/known/durationpb
google.golang.org/protobuf/types/known/timestamppb
google.golang.org/protobuf/types/descriptorpb
github.com/golang/protobuf/ptypes/duration
github.com/golang/protobuf/ptypes/timestamp
github.com/golang/protobuf/ptypes/any
google.golang.org/protobuf/reflect/protodesc
github.com/golang/protobuf/proto
google.golang.org/genproto/googleapis/rpc/status
github.com/golang/protobuf/ptypes
google.golang.org/grpc/internal/status
google.golang.org/grpc/status
github.com/containerd/containerd/errdefs
github.com/containerd/containerd/platforms
github.com/docker/docker/errdefs
github.com/docker/docker/client
go get: added github.com/Microsoft/go-winio v0.5.0
go get: added github.com/containerd/containerd v1.5.2
go get: added github.com/docker/docker v20.10.7+incompatible
go get: added github.com/docker/go-connections v0.4.0
go get: added github.com/sirupsen/logrus v1.8.1
go get: upgraded golang.org/x/net v0.0.0-20210510120150-4163338589ed => v0.0.0-20210525063256-abc453219eb5
go get: upgraded golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 => v0.0.0-20210603125802-9665404d3644
go get: added google.golang.org/genproto v0.0.0-20210607140030-00d4fb20b1ae
```

https://godoc.org/github.com/docker/docker/client

https://pkg.go.dev/github.com/docker/docker/client?utm_source=godoc

Docker Engine API v1.41 - https://docs.docker.com/engine/api/v1.41/ which is the latest now

Latest URL is also - https://docs.docker.com/engine/api/late

Version history - https://docs.docker.com/engine/api/version-history/

Overview - https://docs.docker.com/engine/api/ , https://docs.docker.com/engine/api/#versioned-api-and-sdk

Examples!! https://docs.docker.com/engine/api/sdk/examples/

I'll start with the quick start - https://docs.docker.com/engine/api/sdk/#sdk-and-api-quickstart

```bash
$ go run pkg/platforms/docker/docker.go 
{"status":"Pulling from library/alpine","id":"latest"}
{"status":"Pulling fs layer","progressDetail":{},"id":"540db60ca938"}
{"status":"Downloading","progressDetail":{"current":29354,"total":2811969},"progress":"[\u003e                                                  ]  29.35kB/2.812MB","id":"540db60ca938"}
{"status":"Downloading","progressDetail":{"current":735671,"total":2811969},"progress":"[=============\u003e                                     ]  735.7kB/2.812MB","id":"540db60ca938"}
{"status":"Downloading","progressDetail":{"current":1649079,"total":2811969},"progress":"[=============================\u003e                     ]  1.649MB/2.812MB","id":"540db60ca938"}
{"status":"Downloading","progressDetail":{"current":2632119,"total":2811969},"progress":"[==============================================\u003e    ]  2.632MB/2.812MB","id":"540db60ca938"}
{"status":"Verifying Checksum","progressDetail":{},"id":"540db60ca938"}
{"status":"Download complete","progressDetail":{},"id":"540db60ca938"}
{"status":"Extracting","progressDetail":{"current":32768,"total":2811969},"progress":"[\u003e                                                  ]  32.77kB/2.812MB","id":"540db60ca938"}
{"status":"Extracting","progressDetail":{"current":98304,"total":2811969},"progress":"[=\u003e                                                 ]   98.3kB/2.812MB","id":"540db60ca938"}
{"status":"Extracting","progressDetail":{"current":1277952,"total":2811969},"progress":"[======================\u003e                            ]  1.278MB/2.812MB","id":"540db60ca938"}
{"status":"Extracting","progressDetail":{"current":2719744,"total":2811969},"progress":"[================================================\u003e  ]   2.72MB/2.812MB","id":"540db60ca938"}
{"status":"Extracting","progressDetail":{"current":2811969,"total":2811969},"progress":"[==================================================\u003e]  2.812MB/2.812MB","id":"540db60ca938"}
{"status":"Pull complete","progressDetail":{},"id":"540db60ca938"}
{"status":"Digest: sha256:69e70a79f2d41ab5d637de98c1e0b055206ba40a8145e7bddb55ccc04e13cf8f"}
{"status":"Status: Downloaded newer image for alpine:latest"}
hello world
```

That's so cool!! :D

It's interesting to see the code! :)

Looking at the code for running Docker container, it's interesting!

I always use docker run, as it's easy to run a container like that

I do remember seeing create and start commands, not to mention stop, rm (remove) etc. I actually never stopped containers later. At some point I always started to force remove containers which stopped and removed even running containers, instead of giving errors that running containers cannot be removed etc

In the code, we clearly are doing the following steps -

We ensure that the container image is present in the Docker Engine by pulling the image

We create a container

We then start the container

We wait for the container to stop

We get the container logs

We actually also get logs for the pull ! :)

So, run is basically - create a container and then start it

Now, for the API - I can't keep the client waiting for a long time till the API server talks to Docker engine, pulls images, create a container, start it, and get logs etc

The API is to run a container. It can be an asynchronous operation - so, we can let the API server send back the response immediately after it captures the operation - "run a container"

I think we will have to think about server, worker model. As it doesn't make sense for the server to take care of talking to the platform too and doing asynchronous tasks. But anyways, for now I'll let the server codebase itself do it, and see how it can be better managed with a separate worker process, running as a separate binary and maybe even in a different machine, to enable a distributed system. We might need some queues for such stuff


Server -> Job Queue -> Worker

Worker would listen and look for Jobs in the Job Queue and pick up the Job if one is present. We would also need to understand how to provide the Job status - container run failed etc. Probably more of a futuristic thing. Anyways. For now, every operation can be a separate job - run container, get status, get logs etc

Also, there could be more than one worker. If one worker picks up a job, it should not be picked up by another. If a worker fails to run a job due to some worker issue, then another worker can pick it up and run again. But those are some extreme cases

For now, I plan to have an in-memory job queue using channels maybe and an in-memory worker module as part of the api server binary :)

---

I didn't implement any worker as of now, and the API call looks like this -

```bash
$ go run server.go
```

```bash
$ curl localhost:8080
Started container
```

```
$ docker ps -a
CONTAINER ID   IMAGE                                     COMMAND                  CREATED         STATUS                     PORTS     NAMES
32aa06a168c1   alpine                                    "echo 'hello world'"     4 seconds ago   Exited (0) 3 seconds ago             admiring_kilby
842d0a03626b   alpine                                    "echo 'hello world'"     2 hours ago     Exited (0) 2 hours ago               agitated_darwin
```

```bash
$ time curl localhost:8080
Started container!
real	0m4.135s
user	0m0.004s
sys	0m0.005s
```

That's a total of 4 seconds!! Wow! Hmm

```bash
$ time curl localhost:8080
Started container!
real	0m4.095s
user	0m0.004s
sys	0m0.006s

$ docker ps -a
CONTAINER ID   IMAGE     COMMAND                CREATED         STATUS                     PORTS     NAMES
c7fb911a6520   alpine    "echo 'hello world'"   4 seconds ago   Exited (0) 3 seconds ago             heuristic_sinoussi
```

Cool right? :D
