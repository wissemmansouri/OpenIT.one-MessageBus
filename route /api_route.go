package route

import (
	"github.com/wissemmansouri/OpenIT.one-MessageBus/codegen"
	"github.com/wissemmansouri/OpenIT.one-MessageBus/service"
	jsoniter "github.com/json-iterator/go"
)

type APIRoute struct {
	services *service.Services
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewAPIRoute(services *service.Services) codegen.ServerInterface {
	return &APIRoute{
		services: services,
	}
}
