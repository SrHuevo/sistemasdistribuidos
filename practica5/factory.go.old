package main

import (
	"fmt"
	"practica5/sem"
	"strconv"
	"time"
	"strings"
)

//Piezas máximas que se construiran de cada componente
const maxPiezas = 100

//Número de robots que contruiran móbiles
const nRobot = 2

//Número de piezas distintas necesarias para contruir un movil
const piezasDist = 4

//Piezas que necesitará el robot de cada cosa y posición que ocuparán
//en el array *Sumar a piezasDist si se añaden piezas nuevas
const cable = 0
const cablesNed = 5
const cablesPos = 0
const pantalla = 1
const pantallasNed = 1
const pantallasPos = cablesPos + cablesNed
const carcasa = 2
const carcasasNed = 1
const carcasasPos = pantallasPos + pantallasNed
const placa = 3
const placasNed = 1
const placasPos = carcasasPos + carcasasNed

//Total de piezas necesarias
const piezasNed = placasPos + placasNed

type pieza struct {
	holesem  *sem.Sem
	piezasem *sem.Sem
	ids      []int
	index    int
}

var piezas = make([]*pieza, piezasDist)

func initPiezas() {
	for i := 0; i < piezasDist; i++ {
		hSem := sem.NewSem(maxPiezas)
		pSem := sem.NewSem(0)
		ids := make([]int, maxPiezas)
		piezas[i] = &pieza{hSem, pSem, ids, 0}
	}
}

func (p *pieza) produce() {
	i := 0
	for id := 0; ; id++ {
		p.holesem.Down()
		p.ids[i] = id
		i = (i + 1) % maxPiezas
		p.piezasem.Up()
	}
}

func (p *pieza) get() string {
	p.piezasem.Down()
	id := p.ids[p.index]
	p.index = (p.index + 1) % maxPiezas
	p.holesem.Up()
	return strconv.Itoa(id)
}

func robot(idRobot string) {
	piezasIds := make([]string, piezasNed)
	for i := 0; ; i = (i + 1) % piezasNed {
		switch {
		case i < pantallasPos:
			piezasIds[i] = piezas[cable].get()
		case i < carcasasPos:
			piezasIds[i] = piezas[pantalla].get()
		case i < placasPos:
			piezasIds[i] = piezas[carcasa].get()
		case i < piezasNed:
			piezasIds[i] = piezas[placa].get()
			makeMobile(idRobot, piezasIds)
		}
	}
}

func makeMobile(idRobot string, piezasIds []string) {
	fmt.Println(printIds(idRobot, piezasIds), "Comenzando")
	time.Sleep(200 * time.Millisecond)
	fmt.Println(printIds(idRobot, piezasIds), "Terminado")
}

const nPrints = piezasDist + piezasNed * 2 + 2
func printIds(id string, piezasIds []string) string {
	s := make([]string, nPrints)
	i := 0;
	s[i] = "robot "
	i = i + 1
	s[i] = id
	i = i + 1
	s[i] = ", cables "
	i = i + 1
	for j := cablesPos; j < cablesNed + cablesPos; j++ {
		s[i] = piezasIds[j]
		i = i + 1
		s[i] = " "
		i = i + 1
	}
	s[i] = ", pantalla "
	i = i + 1
	for j := pantallasPos; j < pantallasNed + pantallasPos; j++ {
		s[i] = piezasIds[j]
		i = i + 1
		s[i] = " "
		i = i + 1
	}
	s[i] = ", carcasa "
	i = i + 1
	for j := carcasasPos; j < carcasasNed + carcasasPos; j++ {
		s[i] = piezasIds[j]
		i = i + 1
		s[i] = " "
		i = i + 1
	}
	s[i] = ", placa "
	i = i + 1
	for j := placasPos; j < placasNed + placasPos; j++ {
		s[i] = piezasIds[j]
		i = i + 1
		s[i] = " "
		i = i + 1
	}
	return strings.Join(s, "")
}

func main() {
	initPiezas()
	go piezas[cable].produce()
	go piezas[carcasa].produce()
	go piezas[pantalla].produce()
	go piezas[placa].produce()

	for id := 0; id < nRobot; id++ {
		go robot(strconv.Itoa(id))
	}
	s := sem.NewSem(0)
	s.Down()
}
