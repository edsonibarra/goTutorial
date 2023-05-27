package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Latitud, Longitud float64
	Date time.Time
}

type Point struct {
	X, Y int
}

func main() {
	bootcamp := Bootcamp{Latitud:40.730610, Longitud: -73.935242, Date: time.Now()}
	fmt.Println(bootcamp)

	p1 := Point{X: 100, Y: 200}
	fmt.Println(p1)

	p2 := Point{} // x: 0, y: 0
	fmt.Println(p2)
}