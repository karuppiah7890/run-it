package main

import "github.com/karuppiah7890/run-it/api/pkg/platforms/docker"

func worker(channel <- chan string) {
	for {
		<- channel
		docker.RunContainer()
	}
}
