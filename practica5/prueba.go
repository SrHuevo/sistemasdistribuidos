package main

import (
	"fmt"
)

func main() {
	i := make(map[string]int)
	fmt.Println(i)
	i["hola"] = 1
	fmt.Println(i["hola"])
}
