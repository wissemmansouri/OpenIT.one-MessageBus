package in

import (
	"github.com/wissemmansouri/OpenIT.one-MessageBus/codegen"
	"github.com/wissemmansouri/OpenIT.one-MessageBus/model"
)

func PropertyTypeAdapter(propertyType codegen.PropertyType) model.PropertyType {
	return model.PropertyType{
		Name: propertyType.Name,
	}
}
