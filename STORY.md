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
