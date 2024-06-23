package main

import (
	"fmt"
	"piscine"
)

// Write a function called Enigma that receives pointers as arguments and move its values
// around to hide them. This function will put:
// • a into c
// • c into d
// • d into b
// • b into a

func main() {
	x := 5
	y := &x
	z := &y
	a := &z
	w := 2
	b := &w
	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j
	k := 6
	l := &k
	m := &l
	n := &m
	d := &n
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
	piscine.Enigma(a, b, c, d)
	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
}
