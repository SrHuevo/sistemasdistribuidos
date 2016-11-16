package main

import (
	"fmt"
)

func main() {
	for i:= 0, j:=0;i < 10; i++,j+2 {
		fmt.Println(i, " ", j)
	}
}
