package out

import (
	"time"

	"github.com/wissemmansouri/OpenIT.one-Common/utils"
	"github.com/wissemmansouri/OpenIT.one-MessageBus/codegen"
	"github.com/wissemmansouri/OpenIT.one-MessageBus/model"
)

func EventAdapter(event model.Event) codegen.Event {
	// properties := make([]codegen.Property, 0)
	// for _, property := range event.Properties {
	// 	properties = append(properties, PropertyAdapter(property))
	// }

	return codegen.Event{
		SourceID:   event.SourceID,
		Name:       event.Name,
		Properties: event.Properties,
		Timestamp:  utils.Ptr(time.Unix(event.Timestamp, 0)),
		Uuid:       &event.UUID,
	}
}
