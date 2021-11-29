package main

import (
	"fmt"
	"strings"
)

func main()  {
	s := "asd,asd,asd,sdfg"

	s1 := strings.Split(s,",")
	fmt.Println(s1)


	s2 := strings.Join(s1,",")
	fmt.Println(s2)
}
