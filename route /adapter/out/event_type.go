package out

import (
	"github.com/wissemmansouri/OpenIT.one-MessageBus/codegen"
	"github.com/wissemmansouri/OpenIT.one-MessageBus/model"
)

func EventTypeAdapter(eventType model.EventType) codegen.EventType {
	propertyTypeList := make([]codegen.PropertyType, 0)
	for _, propertyType := range eventType.PropertyTypeList {
		propertyTypeList = append(propertyTypeList, PropertyTypeAdapter(propertyType))
	}

	return codegen.EventType{
		SourceID:         eventType.SourceID,
		Name:             eventType.Name,
		PropertyTypeList: propertyTypeList,
	}
}
