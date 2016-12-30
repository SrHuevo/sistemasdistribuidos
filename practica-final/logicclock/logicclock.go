package logicclock

import (
    "encoding/json"
)

struct Event type {
    MapEvents map[string]int
    Id string
    Log string
}

func (Event e) ToJson()string {
  return string(json.Marshal(e))
}

func (Event e) FromJson(string event) {
  e := json.Unmarsal([]byte(event))
}
