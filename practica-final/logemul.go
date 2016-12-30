package main

import (
	"logiclog"
	"sync"
	"rand"
	"log"
)

var chans = make([]chan int, 6)

func superproceso(log string) {
	fOut, err := os.Create(pathFile)
	if err {
		log.Fatalln("No se ha podido crear el fichero")
	}
	lc := logiclog.Create(fOut, "Proceso 1")
	for i:= 0;; i++ {
		if(rand.Int() > 0.1){
			logiclog.Log()
		}
		if(rand.Int() > 0.8){
			logiclog.
		}
	}
}
//prueba
func upProcess(log string, index int) {
	chans[index] = make(chan int)
	go run(log)
}

func main() {
	upProcess("soy el proceso 1",0)
	upProcess("soy el proceso 2",1)
	upProcess("soy el proceso 3",2)
	upProcess("soy el proceso 4",3)
	upProcess("soy el proceso 5",4)
	upProcess("soy el proceso 6",5)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
