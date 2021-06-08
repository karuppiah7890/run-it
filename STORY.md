# Story

This is the story of how run-it idea came up and how it's going to be built

I always loved the online coding platforms and environments like

https://replit.com

https://glitch.com/

https://jsbin.com/

https://codesandbox.io

https://codepen.io/

https://katacoda.com/

https://labs.play-with-docker.com

There are also so many tutorial websites and hacking and interview platforms -

https://freecodecamp.org

https://www.hackerrank.com/

https://www.hackerearth.com/

And there are probably many others. I am not able to recall the others that I used now but clearly these are all cool ! Not to mention the multitude of online services that allow one to write code, build and run them completely online without installing anything on their machine except a browser - that's a pretty cool thing!! :) 

I was wondering how they build these platforms, so I wanted to try something and hack stuff and probably also learn something from this whole experience :)

Some things I have in mind in terms of features in a step by step manner -

- As a user, I should be able to open up a Web UI and type code in a single file and run it and see it's output in a console like view in the Web UI

- As a user, I want to be able to type code in different languages - Golang, Ruby, Java, JavaScript, Python and more!

- As a user, I want to be able to create multiple projects and write code in a single file in the different projects and then run the code in the projects

- As a user, in each project I want to be able to write code in a multiple files and then run the code in the project

There are also some things that I had in mind for the administrator

- As an admin, I want the user to be able to run only one build at a time to avoid building same or different versions of the code multiple times concurrently, as this leads to some builds running older versions of the code while others running newer versions of the code, or in some cases, many builds running same version of the code, all at the same time. This is not a useful way to use the build infrastructure

- As an admin, I should be able to use different platforms to run the user's code - Docker, Podman, CRI-O, containerd, Kubernetes, Nomad, Docker Swarm are some of the platforms I would like support for, starting with support for Docker for MVP

- As an admin, I want the user to be able to run their code in an isolated manner in the infrastructure in order to avoid any security issues

- As an admin, I want the user's code to not use too many resources like CPU, RAM, Storage, which can hog the infrastructure resources and impact the cost of running the infrastructure

- As an admin, I want the user to be able to run

- As an admin, I want a portal to be able to configure the platforms

- As an admin, I want a portal to be able to configure the run-times for languages available for the users so that the users can choose which runtime to use - language + version. For example - Golang v1.16.1, Golang v1.15.1 etc. The configuration should have runtime info - description for users to read like OS, etc and then something like a Docker / OCI image URL which contains the runtime

- As an admin, I want to cancel existing builds of the user automatically and instead run new builds if the code has changed and user wants to build and run the new version of the code by clicking build and run

- As an admin, I want the user to be able to run code in multiple projects at the same time

- As an admin, I want my portal and user's portal to be protected with credentials

- As an admin, I want many users to be able to sign up and login

- As an admin, I want many users to be able to run code at the same time

- As an admin, I want users to be able to setup multi factor authentication using Time based One Time Password (TOTP)


---

Okay, that was a hell a lot of features. Lol. I got carried away at some point. LOL 😆

The goal is to be able to learn a lot of things out of this project. Especially things related to the backend - Docker / OCI API etc, and also Kubernetes, Nomad API etc

Some nice technical things I can learn are -

How to save code that the user types

How to run the code that the user types

How to run the code safely - by not running as root but instead running the code as non-root

How to run the code safely - by restricting the resources - CPU, RAM

How to run the code for a given amount of time and not for too long? As the user can run some server code which can block and run for a long time. So, we need to have a simple program which runs and exits and not a long running process like a server / backend service

Questions
- Should network be restricted too? Hmm. Like, downloading something from the Internet and uploading to some system on the Internet, using up the Internet / network resource of the infrastructure

Some technical thoughts around design of code -

I want to have a web UI and a backend, both separate. Web UI will be a static site

I want to keep the API server code and the platform specific code separately. For example, if I support Docker first, and then later Kubernetes, I want to have all this platform specific code separately :) And then API server code separately, as it kind of makes sense to have this separation of concerns

Requests to API server will look like - save code / run code / give logs of the run etc

And accordingly the API server can integrate with the configured platform and do these actions :)

It will be interesting to see how I get real time logs! :D I'll learn Docker, Kubernetes log API

To start off, I wanted to support only Golang code and only one file code, and only run with Docker

Given that Docker is a container runtime / container engine, I was also wondering about checking out OCI and if it has generic tools, clients and client libraries to interact with OCI compliant container runtimes if that's a thing. In which case, I can write one code and use it for Docker, Podman, CRI-O, containerd, rkt etc - basically any container runtime that's compliant with the OCI spec :D Same for container images :D

How will the API look like? Hmm

API will have this one endpoint to list out the runtimes / environments it has - like Golang v1.16.1, Python 3.7 etc. We can call it runtime. And we don't have to tell the user as to how or where the code is running - I mean, they don't need to know the platform, I think. It can run on Docker, or Kubernetes etc. It doesn't matter, as that's an internal implementation detail / backend detail. For the user, the code has to be run in some environment with a particular version of a language runtime, like Golang v1.16.3

In short, some of the things to do are - Save code, Build code, Run code :)

Also, for any restrictions, for example - initially not running multiple builds at the same time - I need to add restrictions at API server level and not just UI level :) As someone can do a API call if UI blocks things and that can't be allowed!

---

New Feature! Idea! Sometimes people want to use environment variables, we can provide a separate place to store the environment variables ;)

---

New Feature! Idea! People cannot build all kinds of code in this environment. For example, sometimes some code has references to private repo code. For example, in golang, one can refer to private repos in `import` statements. In such cases, people need access to those git repos in their local to make it work. In this code environment, that's not possible! As the code environment will not have credentials or extra access to pull the private code, also that's too much work, depending on the kind of language and module system etc. So, we can provide a FAQ about what's possible and what's not possible. Especially what's not possible. The admin can provide an FAQ for each runtime or for all runtimes in general in the form of a JSON - array of objects containing `question` and `answer` fields which the front end can pull from the API server and can render when the user wants to see FAQ for the specific runtime or for all runtimes in general. It can be merged together by the UI and shown to the user actually

---

New Feature! Idea! About resources - we need to ensure that the source code is not too big, for now at least. For which we can check size, or do something else too. For example, we can store the source code in platform specific storage - for Docker, maybe Docker volumes? For Kubernetes, PV? The storage can be user specific - containing multiple projects, or project specific where one project has one storage. Later when users share code, it will be easy to have one storage for one project, like one Docker volume for one project.

---

New Feature! Idea! Users can fork code from other users

---

New Feature! Idea! Users can make their code public or private ;) Public code can be seen by other users through some gallery view maybe? :D

---

New Feature! Idea! Version the user's code? This is a bit much though. We will need some system like Git :P What's the purpose? Can user go back to an older version? How is commit / save time decided? By the user? By the system? Hmm

---

For saving code or mounting code in an environment to run the code, I was wondering how it can be done

For starters, I thought I can do volume mount when using Docker. I was thinking how volume mount works - I just realized that it might be transferring files from the client to the server. Gotta confirm though. Anyways, for now, the code is just one file and we can restrict the code size to avoid resource hogging of the storage resource

Later, we can think about how to store code in something like Docker volume and attach the Docker volume to a container

Similarly, we need to see how to isolate the network(s) of the containers. Gotta ensure it's by default isolated, or else isolate it, so that one container cannot connect to other containers by being in the same network and also accessible :)

For golang, I can use `go run` instead of doing `go build` and then using some binary name. So, it's gonna be

```bash
$ go run main.go
```

Instead of

```bash
$ go build -v -o app
$ ./app
```

The user should also be able to understand that the environment is a linux environment and what OS it is / has etc, for example Ubuntu, Arch Linux etc. This has to be provided by the admin when configuring the runtime. For example, it should be like -

Runtime name - Go v1.16.3
Runtime code - go-v1.16.3
OS: Ubuntu 20.04
Arch: amd64

Something like that. I think the arch also depends on the platform, like, Docker engine amd64 can only run amd64 container images I think. Something like that. Gotta confirm those stuff

Why is this important? Sometimes users might be writing platform specific code in which case the build and run might fail ! Such platform specific code is very minimal though :) At least in golang. Not sure how platform specific stuff is handled in other languages, but there might be some similar way using just code :)

---

Next steps?

Try writing golang code in a single file

Try running it in a Docker container using Docker CLI

Try running it in a Docker container using OCI tools and CLI

Try running it in a Docker container using OCI using Golang code and libraries

---

Some decisions - I plan to use Golang to write the code for the platform. Let's see how that goes :)

I plan to use some data store for storage. It might a bit complex like PostgreSQL or something a bit simple like Redis, etcd. I know only these stores in general. I gotta think about what kind of data I want to store and then think about what Database to use

---

Running with golang Docker image from Docker Hub

https://hub.docker.com/_/golang

```bash
$ docker run --rm -v experiments/golang-and-docker/main.go:/app/main.go  golang:1.16.5 go run /app/main.goUnable to find image 'golang:1.16.5' locally
1.16.5: Pulling from library/golang
d960726af2be: Already exists 
e8d62473a22d: Already exists 
8962bc0fad55: Already exists 
65d943ee54c1: Already exists 
f2253e6fbefa: Already exists 
186c77a2a533: Pull complete 
db807893dccf: Pull complete 
Digest: sha256:360bc82ac2b24e9ab6e5867eebac780920b92175bb2e9e1952dce15571699baa
Status: Downloaded newer image for golang:1.16.5
docker: Error response from daemon: create experiments/golang-and-docker/main.go: "experiments/golang-and-docker/main.go" includes invalid characters for a local volume name, only "[a-zA-Z0-9][a-zA-Z0-9_.-]" are allowed. If you intended to pass a host directory, use absolute path.
See 'docker run --help'.
```

Oops, syntax issues, haha. It thought the input was a Docker volume name, but it's actually a local file in a local directory, let's fix that!

```bash
$ docker run --rm -v $(pwd)/experiments/golang-and-docker/main.go:/app/main.go  golang:1.16.5 go run /app/main.go
Hello world! :D

$ go run experiments/golang-and-docker/main.go 
Hello world! :D
```

Woohoo! :D Simple thing, done! :)

Now, let's try some CRI tools. Oops, actually OCI tools. Or, OCI CRI I guess :P

Let's learn!

https://opencontainers.org/

https://opencontainers.org/about/overview/

https://github.com/opencontainers

https://github.com/opencontainers/runc

https://github.com/opencontainers/image-tools

https://github.com/opencontainers/runtime-tools

Let's try out runc and see if it helps in interacting with Docker runtime

https://github.com/opencontainers/runc/releases

I'm forgetting that containers usually need the OS features to do containerization - isolation of processes, or running isolated processes. chroot, namespaces, cgroups. Linux has it and I think other OSes are trying to bring those features, but I don't know much. I remember some mention of newer versions of Windows Server bringing support for containerization features similar to virtualization

A mention in runc README

https://github.com/opencontainers/runc

```
runc currently supports the Linux platform with various architecture support. It must be built with Go version 1.13 or higher.

In order to enable seccomp support you will need to install libseccomp on your platform.
```

In my case, my machine is Mac. Hmm. I'm wondering how I can run runc and connect runc with the Docker engine running inside some Linux VM inside my Mac. I first need a mac based runc. Hmm. And I don't think I can build that - if there's OS specific code in it, for example Linux specific. Hmm. I guess I gotta try. Or, I could just first write code with Docker client and then check more about this. Or, I could check if there are OCI client libraries to interact with multiple container runtimes the same way!! :) :D

I was looking for more docs

https://github.com/opencontainers/runc/tree/master/docs

https://github.com/opencontainers/runtime-tools/tree/master/docs

https://github.com/opencontainers/runtime-tools/blob/master/docs/command-line-interface.md

I was also looking for client libraries

https://github.com/search?utf8=%E2%9C%93&q=oci%20client%20golang%20library

https://github.com/search?q=oci+client+library

https://github.com/search?q=oci+client

I guess oci runtime code is more for container runtimes to use, but looks like I can't use the oci runtime code to interact with container runtimes in a polymorphic manner

Here's again a doc in runc README

```
Please note that runc is a low level tool not designed with an end user in mind. It is mostly employed by other higher level container software.

Therefore, unless there is some specific use case that prevents the use of tools like Docker or Podman, it is not recommended to use runc directly.

If you still want to use runc, here's how.
```

So, for now I'll just look for clients to be able to run a Docker client as a golang library

https://github.com/search?q=docker+client

https://github.com/search?l=Go&q=docker+client&type=Repositories

https://github.com/fsouza/go-dockerclient

https://github.com/docker/engine-api

https://github.com/docker/docker/tree/master/client

https://github.com/moby/moby/tree/master/client

I was wondering which one to use. 

https://github.com/fsouza/go-dockerclient or official one

I see this -

https://github.com/fsouza/go-dockerclient#difference-between-go-dockerclient-and-the-official-sdk

```
For new projects, using the official SDK is probably more appropriate as go-dockerclient lags behind the official SDK. 


When using the official SDK, keep in mind that because of how the its dependencies are organized, you may need some extra steps in order to be able to import it in your projects (see #784 and moby/moby#28269).
```

So, I guess I can use the official SDK! :)

https://docs.docker.com/engine/api/sdk/

---

Some more Ideas!! Random ones too :P

For run-it, should I store every run by the user and show it as part of old runs? Where they can see that version of the code and also the run log. In which case all logs need to be stored, hmm. Or there can be a limit. Show last five runs etc. Archive or throw others etc. Hmm. Users cannot edit the code while seeing old run or run that code. That's old code. But they can copy it etc. Hmm. Something like that. As latest code might be different and this would be older version of the code

For resources, maybe admin can allocate max amount of resources for each user. User can then use that max amount of resources for their programs! So, they can run multiple projects simultaneously with some resources for each project :) once a build is raised, the resources used by that build is locked and not usable for other builds until the build is done or stopped and then those resources are available in the user's resource pool. Cool, right? :D but hard too :p keeping track of resources etc

I mentioned stop build. So...what if a build takes a long time? Can the user stop the build? Or when the code is running, can the user stop it in between?

"Building..."

"Running..."

All those notifications / information in the Web UI, not in the console like UI, but in the web UI. And user clicks "stop running" or "stop building"? Well, you stop it. The infrastructure should stop it, remove container, etc :) user may have their own reasons to stop the build and or run, hmm. Especially if user's resources are shared across projects, then they could stop the build or run, to use the resources in another project :) :D

---

Some more ideas!!

Logs - how to design an API to get logs? Will it be like "get all logs for the run"? Or will it be paginated?

For example, if I get the logs like

/logs?n=5

And get only first line of log instead of 5 lines since the container is just getting started, then, I can't send request next to get the next 5, because I haven't got the first five itself. I actually need to send request to get from second one. So, it could be paginated like

/logs?from=1&to=5

And if it gives only one line, it can also return next reference

next=2

Or something like that. Maybe the size can always be same? 🤷‍♂️ Always five lines. Or instead of "to" we can say log size

size=5

And use a default value when it's not provided, like 5 or 10 etc

I was also thinking if I should get full container logs maybe, instead of paginated. But that seemed too much. Too much data. But that's also kind of an easy implementation

It all actually depends on how we can get logs from the platforms too :) gotta see how the platforms provide an API for logs and how the response looks like. For example Docker engine API, Kubernetes API, Nomad API etc

I mean, idk how the response looks like or even the request. One response I can think of is - to differentiate different lines, use an array of strings. Or use "\n" and have one big string as response. I don't know

Also, I gotta check how to show all logs, including blank lines which are tricky. What if a program shows multiple blank lines and then puts some data. I need to show the same in the UI and in the API. I need to be able to differentiate between blank line log vs no log. Maybe "\n" will be the differentiator, like, how we put it, and maybe empty logs like "   " etc where just spaces are there or nothing is there, just empty string, and then new line for next log line

Also, for web UI, instead of polling, I can also think of using server sent events to get logs from the API server. Something to think about

---

I was also thinking about monitoring, logging, instrumentation, error monitoring for the API server itself

For example, for basic monitoring, we can have a Prometheus metrics endpoint and export metrics like platform metrics depending on the platform. Like, how many runs have been done with the platform, how many containers are running currently and stuff like that. Same for other platforms

Metrics for platforms would look like docker_total_runs or something like that. We need the information that it's Docker platform. We can also prefix platform in it, if needed and we can export metrics for all platforms. Anything that's active will show up values while others give 0 values or nil / null values in case they are not being used.

As of now the plan is to use only one platform at a time in an API server. In the future maybe this could be changed. I think this might change :) but I need to decide how the platform will be chosen, if user is abstracted away from that detail, and if only admins know it. Maybe something like, a user can use only a particular platform, or a user can use any of the platforms and API server chooses randomly or based on some mechanism, like availability, success rate etc?

Also, for metrics, if we have metrics, we can also have dashboards using Grafana. Like, a graph showing number of containers run in Docjer in the last few days. And a single counter showing total containers run in Docker. And any metric we can think of ! :)

Come to think of it, this is like trying to run CI/CD tasks or jobs and checking it's logs. I was literally thinking of how Travis CI would show it's long list of logs. Maybe I could check their API response! Same for other CI CD systems like GitHub Actions, GitLab etc. For some I can even find the code as it's Open Source. Like Jenkins, Tekton etc :)

I was thinking about metrics because it's important to know what's going on in the system. And monitoring is key for that

We would also need error monitoring to monitor errors. For example Sentry like systems

I'm talking about this because it's important to understand what's going on in the system when I'm developing and running the system in development mode and in a production environment

Ideally it shouldn't be too hard to debug in development mode. Maybe I can log a ton of things during development by keeping log level as too high. I'm just wondering if there are better ways to understand what's going on in the system, hmm

---

Also, I thought about the UI, for now I think building, running, it's all the same? Atleast for Golang, maybe for Java we have to build with Java compiler (JDK), and then run with JRE. Even with Golang that can be done, compilation and then running, but I'm going with a simple method for now

So, for now, I'm just gonna be showing as "running"

And the UI will have two buttons - play for run, and a square for stopping. Play will be in green. Stop will be in red. I can also put text in the buttons and on hover and anywhere needed so that it's accessible :) color blind people, people with other eye disabilities would be able to use the text and things like aria labels to understand what the buttons do! :)

---

Though I have so much ideas, for UI etc, I wanna start with the API and keep the API clean and very good in terms of design :)

And then there's UI, monitoring, logging etc

I might do some dirty logging for now, a lot of it maybe, for development purposes

---

I'm going to start off by creating an API to run some golang code with Docker container! :) The story is [here](./stories/api-to-run-with-docker-container.md)

---

Ideas for running in Kubernetes

For running code in Kubernetes, I could use Jobs, or I could use pods directly too. If I use pods directly, there will not be any retry if it fails, if I set some options especially. Since it's more of a task or process taht runs to completion. Or else Kubernetes might think it died and might restart it etc. I gotta ensure it is in a normal and expected state, avoiding multiple runs etc

Jobs can help in case of failures and retries or even parallelism etc, but that's not necessary here 🤷‍♂️

I also don't need complicated stuff like deployment. I'm not really deploying a service or anything. No updates, no multiple instances etc. So, just pod is good I think

---

Ideas for Docker container timeout or max time, probably in the future

```bash
$ docker run --help | rg time
      --cpu-rt-period int              Limit CPU real-time period in microseconds
      --cpu-rt-runtime int             Limit CPU real-time runtime in microseconds
      --health-timeout duration        Maximum time to allow one check to run (ms|s|m|h) (default 0s)
      --runtime string                 Runtime to use for this container
      --stop-timeout int               Timeout (in seconds) to stop a container

$ docker run --help | rg stop
      --stop-signal string             Signal to stop a container (default "SIGTERM")
      --stop-timeout int               Timeout (in seconds) to stop a container
```

But the timeout didn't work ! :/ I'm not sure why

```bash
$ time docker run --rm --stop-timeout 5  golang:1.16.5 sleep 20

real	0m20.742s
user	0m0.142s
sys	0m0.117s

$ docker run --help | less

$ sleep 20

$ sleep 50
Terminated: 15

$ kill --help

$ docker run --help | less

$ time docker run --rm --stop-timeout aa golang:1.16.5 sleep 20
invalid argument "aa" for "--stop-timeout" flag: strconv.ParseInt: parsing "aa": invalid syntax
See 'docker run --help'.

real	0m1.035s
user	0m0.216s
sys	0m0.398s
$ time docker run --rm --stop-timeout 1 golang:1.16.5 sleep 20

real	0m20.749s
user	0m0.145s
sys	0m0.122s

$ time docker run --rm --stop-timeout 1 --stop-signal SIGKILL golang:1.16.5 sleep 20

real	0m20.719s
user	0m0.144s
sys	0m0.114s

$ time docker run --rm --stop-timeout 1 --stop-signal KILL golang:1.16.5 sleep 20
^C
real	0m20.745s
user	0m0.145s
sys	0m0.115s
```

https://duckduckgo.com/?t=ffab&q=docker+stop+timeout&ia=web

https://stackoverflow.com/questions/48299352/how-to-limit-docker-run-execution-time#48299490

```bash
$ time docker run --rm --stop-timeout 1 --stop-signal KILL golang:1.16.5 bash -c '(while true; do true; done)'
^C
real	0m16.205s
user	0m0.144s
sys	0m0.118s
$ time docker run --rm --stop-timeout 1 --stop-signal SIGKILL golang:1.16.5 bash -c '(while true; do true; done)'

real	0m12.907s
user	0m0.144s
sys	0m0.072s
```

Interesting, there's an issue for it in Docker core which is moby

https://github.com/moby/moby/issues/1905

https://stackoverflow.com/questions/28933925/docker-timeout-for-container

I think the best way is to simply have a timer in the API server / core module, which will call stop on the platform's resources - containers in the case of Docker platform, pods or something like that in Kubernetes platform and so on

Anyways, this is an interesting thing, hmm
