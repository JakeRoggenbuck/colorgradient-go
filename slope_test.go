package main

import (
	"fmt"
	"testing"
)


func Testfind_slope(t *testing.T) {
	p1_1 := Point {2, 3}
	p2_1 := Point {5, 6}
	slope_1 := find_slope(&p1_1, &p2_1)

	p1_2 := Point {4, 3}
	p2_2 := Point {8, 6}
	slope_2 := find_slope(&p1_2, &p2_2)

	fmt.Println(slope_1)
	fmt.Println(slope_2)
}

func main_test() {
	
}
