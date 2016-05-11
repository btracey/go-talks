package main 

import "fmt"

func main(){
	
	var array [10]int
	var slice []int
	list := []string{"a", "b", "c", "d", "e", "f"}
	other :=append(list[0:2], list[3:5]...)
	for k, v := range other{
		fmt.Printf("%d -> %s\n", k, v)
	}
}