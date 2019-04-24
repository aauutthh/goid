package main

import (
	"fmt"
	"time"

	"github.com/aauutthh/goid"
)

func main() {
	fmt.Println("main goid:", goid.GetGoID())
	fmt.Println("main goid:", goid.GetGoIDAsm())
	go func() {
		fmt.Println("goid:", goid.GetGoID())
		fmt.Println("goid:", goid.GetGoIDAsm())
	}()
	time.Sleep(time.Second * 2)
}
