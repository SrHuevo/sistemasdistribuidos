package main

import (
	"log"
	"os"
	"sistemasdistribuidos/practica-final/logiclog"
	"strconv"
	"sync"
	"time"
)

type recepType struct {
	salaespera []mRecep
	total      int
	indexIn    int
	indexOut   int
}

type mClient struct {
	freePlace bool
	slog      string
}

type mRecep struct {
	slog    string
	cClient chan mClient
}

const maxClients = 5

func initLog(name string) *logiclog.Llog {
	os.Mkdir("logs", 0777)
	f, err := os.Create("logs/" + name + ".txt")
	if err != nil {
		log.Fatal("No se ha podido crear el fichero")
		panic(err)
	}
	return logiclog.Create(f, name)
}

func cliente(r chan mRecep, nc int, llog *logiclog.Llog) {
	c := make(chan mClient)
	r <- mRecep{llog.Mark(), c}
	mc := <-c
	llog.Msg(mc.slog)
	if mc.freePlace {
		llog.Log("Cliente" + strconv.Itoa(nc) + ": me siento en la sala de espera")
		mc = <-c
		llog.Msg(mc.slog)
		llog.Log("Cliente" + strconv.Itoa(nc) + ": me corto el pelo")
		mc = <-c
		llog.Msg(mc.slog)
		llog.Log("Cliente" + strconv.Itoa(nc) + ": termino de cortarme el pelo")
	} else {
		llog.Log("Cliente" + strconv.Itoa(nc) + ": me voy de la barberia, esta llena")
	}
}

func clientes(r chan mRecep) {
	llog := initLog("GeneradorClientes")
	for nc := 1; ; nc++ {
		go cliente(r, nc, llog)
		time.Sleep(1000 * time.Millisecond)
	}
}

func recep(r chan mRecep, b chan mRecep) {
	llog := initLog("recepcion")
	rtype := &recepType{}
	rtype.salaespera = make([]mRecep, 5)
	rtype.total = 0
	rtype.indexIn = 0
	rtype.indexOut = 0
	for {
		if rtype.total == 0 {
			mr := <-r
			llog.Msg(mr.slog)
			enviarASala(mr, rtype, llog)
		} else {
			atiendeClientes(rtype, r, b, llog)
		}
	}
}

func atiendeClientes(rtype *recepType, r chan mRecep, b chan mRecep, llog *logiclog.Llog) {
	select {
	case b <- rtype.salaespera[rtype.indexOut]:
		enviarABarbero(rtype)
	case mr := <-r:
		llog.Msg(mr.slog)
		if rtype.total == maxClients {
			enviarACalle(mr.cClient, llog)
		} else {
			enviarASala(mr, rtype, llog)
		}
	}
}

func enviarABarbero(rtype *recepType) {
	rtype.indexOut = (rtype.indexOut + 1) % maxClients
	rtype.total--
}

func enviarASala(mr mRecep, rtype *recepType, llog *logiclog.Llog) {
	mc := mClient{true, llog.Mark()}
	mr.cClient <- mc
	rtype.salaespera[rtype.indexIn] = mr
	rtype.indexIn = (rtype.indexIn + 1) % maxClients
	rtype.total++
}

func enviarACalle(c chan mClient, llog *logiclog.Llog) {
	mc := mClient{false, llog.Mark()}
	c <- mc
}

func barbero(nb int, b chan mRecep, name string) {
	llog := initLog(name)
	for mr := range b {
		llog.Msg(mr.slog)
		mc := mClient{true, llog.Mark()}
		mr.cClient <- mc
		llog.Log("Empiezo a cortar el pelo")
		time.Sleep(2000 * time.Millisecond)
		llog.Log("Termino de cortar el pelo")
		mc = mClient{true, llog.Mark()}
		mr.cClient <- mc
		llog.Log("Me duermo esperando clientes")
	}
}

func main() {
	b := make(chan mRecep)
	r := make(chan mRecep)

	go barbero(1, b, "barbero1")
	go barbero(2, b, "barbero2")
	time.Sleep(100 * time.Millisecond)

	go recep(r, b)
	time.Sleep(100 * time.Millisecond)

	go clientes(r)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
