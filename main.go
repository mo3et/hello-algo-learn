package main

import (
	"fmt"
)

func main() {
	// var a []int
	// var acf string = "acfun"
	// fmt.Println(a)

	// var ac [5]int
	// fmt.Println(reflect.TypeOf(ac))

	// b := [...]int{1, 2, 3, 0, 0, 0, 1}
	// s1 := make([]int, 3)
	// s2 := append(s1, 233)
	// fmt.Printf(" values is %s %v", acf, s1)
	// fmt.Println(reflect.TypeOf(s2))
	// fmt.Println(len(b), "values is ", b)
	// fmt.Println(s1, "is a slice")

	// slice := []int{1, 2, 3, 4, 5}
	// slice2 := []int{1, 2, 3, 4, 5}
	// mymap := map[string]int{"apple": 1, "banana": 2, "cherry": 3}

	sl1 := []interface{}{3378, 45, 818, "3q"}
	sl2 := make([]string, len(sl1))
	// copy(sl2, sl1)
	for i, v := range sl1 {
		sl2[i] = fmt.Sprintf("%v", v)
	}
	fmt.Println(sl2)
}
