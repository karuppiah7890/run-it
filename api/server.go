package main

import (
	"fmt"

	"github.com/karuppiah7890/run-it/api/pkg/platforms/docker"
	"github.com/valyala/fasthttp"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	docker.RunContainer()
	fmt.Fprintf(ctx, "Started container!")
}

func main() {
	fasthttp.ListenAndServe(":8080", fastHTTPHandler)
}
