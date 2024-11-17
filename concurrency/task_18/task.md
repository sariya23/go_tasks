===========================================================
Задача 18
Что выведет код и почему?
===========================================================
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	ch := 0
	go func() {
		ch = 1
	}()
	for ch == 0 {
	}
	fmt.Println("finish")
}
