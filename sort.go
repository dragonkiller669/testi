package main

import (
	"fmt"
	"sort"
)

func main(){

//	now := time.Now()
//	a := time.Date(
//		2019, 11, 17, 20, 34, 58, 666, time.UTC)
//	b := time.Date(2018, 12, 10, 23, 59, 23, 666, time.UTC)
//	c := time.Date(2017, 12, 2, 9, 53, 20, 666, time.UTC)
//	times := []time.Time{a, b, c}
//	sort.Slice(times)
//	fmt.Println("Times:", times)
//	fmt.Println(now,"\n", a,"\n", b, "\n", c, "\n")
	strinks := []string{"jopa hyttynen,","lahna,","apina,","ampiainen,","ovat tärkeitä ekosysteemeille"}
	sort.Strings(strinks)
//	numero := []int{1, 6, 5}
//	sort.Ints(numero)

	fmt.Println(strinks)
}
