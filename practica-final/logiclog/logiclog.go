package logiclog

import (
  "../logicclock"
  "fmt"
  "os"
)
struct Llog type {
  fOut os.File;
  lastClock Logicclock
}

func Create(pathFile string, id string) Llog {
  me := Llog{}
  me.lastClock.Id := id
  me.lastClock.MapEvents[id]:= 0
  return me
}

func (me Llog) Log(log string) {
  me.lastClock.MapEvents[id]++
  me.lastClock.Log := log
  me.fOut.WriteString(me.LastClock.ToJson)
}

func (me Llog) Llog Mark() string {
  me.lastClock.MapEvents[id]++
  return me.LastClock.ToJson
}

func (me Llog)Llog Msg(json string) {
  me.lastClock.FromJson(json)
}
