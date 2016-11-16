package factory

import (
	"fmt"
	"strings"
	"time"
)

var nPrints int

func Trabaja(idRobot string, ps *Piezas) {
	nPrints = len(ps.cant)*2 + ps.total*2 + 2
	for {
		pids := getPiezas(ps)
		makeMobile(idRobot, pids, ps)
	}
}

func getPiezas(ps *Piezas) []string {
	piezasIds := make([]string, ps.total)
	cant := 0
	for i, p := range ps.p {
		for j := 0; j < ps.cant[i]; j++ {
			piezasIds[cant] = p.get()
			cant++
		}
	}
	return piezasIds
}

func makeMobile(idRobot string, piezasIds []string, ps *Piezas) {
	fmt.Println(printIds(idRobot, piezasIds, ps), "Comenzando")
	time.Sleep(200 * time.Millisecond)
	fmt.Println(printIds(idRobot, piezasIds, ps), "Terminado")
}

func printIds(id string, piezasIds []string, ps *Piezas) string {
	s := make([]string, nPrints)
	i := 0
	cant := 0
	s[i] = "robot "
	i++
	s[i] = id
	i++
	for j, n := range ps.names {
		s[i] = ", "
		i++
		s[i] = n
		i++
		for k := 0; k < ps.cant[j]; k++ {
			s[i] = " "
			i++
			s[i] = piezasIds[cant]
			i++
			cant++
		}
	}
	return strings.Join(s, "")
}
