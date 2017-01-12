package logiclog

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func initLog(name string) *Llog {
	os.Mkdir("logs", 0777)
	f, err := os.Create("logs/" + name + ".txt")
	if err != nil {
		log.Fatal("No se ha podido crear el fichero")
		panic(err)
	}
	return Create(f, name)
}

func TestLogs(t *testing.T) {
	llog := initLog("testlogs")

	llog.Log("log 1")
	llog.Log("log 2")
	llog.Log("log 3")
	llog.Log("log 4")
	llog.Log("log 5")
}

func TestLogs2(t *testing.T) {
	llog1 := initLog("testlogs2.1")
	llog2 := initLog("testlogs2.2")

	fmt.Println(llog1)
	fmt.Println(llog1)
	fmt.Println(llog2)
	s := llog1.Mark()
	fmt.Println(llog1)
	llog2.Msg(s)
	fmt.Println(llog2)
	fmt.Println(llog2)
	fmt.Println(llog1)
}
