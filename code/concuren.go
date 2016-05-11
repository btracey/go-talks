package main

import "fmt"
import "time"

func Calc(i int) {
	fmt.Printf("I'am routine %d just for calcul\n", i)
	time.Sleep(time.Second * i)
	fmt.Printf("%d termined\n", i)
}

func main() {
	for i := 1; i < 5; i++ {
		go Calc(i)
	}
}
