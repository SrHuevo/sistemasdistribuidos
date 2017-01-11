package logiclog

import (
	"log"
	"os"
	"sistemasdistribuidos/practica-final/logicclock"
)

type Llog struct {
	fOut      *os.File
	lastClock logicclock.Event
}

func Create(file *os.File, id string) *Llog {
	return &Llog{file, logicclock.NewEvent(id)}
}

func (me *Llog) Log(slog string) {
	me.lastClock.Add()
	me.lastClock.Log = slog
	s, err := logicclock.ToJson(me.lastClock)
	if err != nil {
		log.Fatal(err)
	}
	me.fOut.WriteString(s)
	me.fOut.WriteString("\n")
}

func (me *Llog) Mark() string {
	me.lastClock.Add()
	slog, err := logicclock.ToJson(me.lastClock)
	if err != nil {
		log.Fatal(err)
	}
	return slog
}

func (me *Llog) Msg(json string) {
	e, err := logicclock.FromJson(json)
	if err != nil {
		log.Fatal(err)
	}
	e.Put(me.lastClock.Get())
	e.Add()
	me.lastClock.Set(e)
}
