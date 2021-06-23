package journal

import (
	"encoding/json"
	"time"
)

type Cost int

type Event struct {
	Timestamp time.Time `json:"timestamp"`
	Event     string    `json:"event"`
}

type GenericEvent struct {
	Event
	KeyValue map[string]interface{}
}

type _GenericEvent GenericEvent

func (e *GenericEvent) UnmarshalJSON(bs []byte) (err error) {
	foo := _GenericEvent{}

	if err = json.Unmarshal(bs, &foo); err == nil {
		*e = GenericEvent(foo)
	}

	m := make(map[string]interface{})

	if err = json.Unmarshal(bs, &m); err == nil {
		delete(m, "timestamp")
		delete(m, "event")
		e.KeyValue = m
	}
	return err
}

type LocalizedValue struct {
	Name          string `json:"Name"`
	NameLocalised string `json:"Name_Localised,omitempty"`
	Count         int    `json:"Count"`
}
