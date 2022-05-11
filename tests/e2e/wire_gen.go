// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package e2e

import (
	"github.com/douyu/jupiter-layout/internal/exampleserver/service"
	"github.com/douyu/jupiter-layout/internal/pkg/grpc"
)

// Injectors from wire.go:

func CreateExampleService() *service.HelloWorld {
	exampleInterface := grpc.NewExample()
	options := service.Options{
		ExampleGrpc: exampleInterface,
	}
	helloWorld := service.NewHelloWorldService(options)
	return helloWorld
}
