package main

import "todotom/go-gin-restful-api/rest"

type Container struct {
	controller rest.Controller
}

func ProvideContainer(controller rest.Controller) Container {
	return Container{controller}
}

func (container Container) Run() {
	container.controller.Run()
}
