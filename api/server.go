package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	channel := make(chan string, 1000)

	go worker(channel)

	fastHTTPHandler := func(ctx *fasthttp.RequestCtx) {
		channel <- "run container"
		fmt.Fprintf(ctx, "Started container!")
	}

	fasthttp.ListenAndServe(":8080", fastHTTPHandler)
}
