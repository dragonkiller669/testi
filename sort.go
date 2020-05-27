
package main

import (
	"fmt"
	"sort"
	s "strings"
)

func main(){


	strinks := []string{"jopa hyttynen,","lahna,","apina,","ampiainen,","ovat tärkeitä ekosysteemeille"}
	sort.Strings(strinks)


	fmt.Println(strinks)
	fmt.Println("Selvennetään vielä: ", s.ToUpper(strinks[4]))
	fmt.Println(strinks[3] + s.Repeat("Jokainen eläin on laulun arvoinen ",3))
}
