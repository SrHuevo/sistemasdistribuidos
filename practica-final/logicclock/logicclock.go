package logicclock

import "encoding/json"

type Event struct {
	Id        string
	MapEvents map[string]int
	Log       string
}

func NewEvent(id string) Event {
	me := make(map[string]int)
	me[id] = 0
	return Event{id, me, ""}
}

func (me Event) Add() {
	me.MapEvents[me.Id]++
}

func (me Event) Put(i int) {
	me.MapEvents[me.Id] = i
}

func (me Event) Get() int {
	return me.MapEvents[me.Id]
}

func (me Event) Set(e Event) {
	for k, v := range e.MapEvents {
		val, ok := me.MapEvents[k]
		if ok && v < val {
			continue
		}
		me.MapEvents[k] = v
	}
}

func ToJson(e Event) (string, error) {
	v, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(v), nil
}

func FromJson(data string) (Event, error) {
	var e Event
	err := json.Unmarshal([]byte(data), &e)
	return e, err
}
