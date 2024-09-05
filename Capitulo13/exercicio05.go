package main


import "fmt"

type square struct {
	perimetro float64
	diagonal float64
	side float64
}

type circle struct {
	raio float64
	circunferencia float64
	diametro float64
}

func (x circle) calculoArea () float64 {
	const pi = 3.1415
	area := pi*x.raio*x.raio
	return float64(area)
}

func (x square) calculoArea () float64 {
	area := x.side * x.side
	return float64(area)
}

type figure interface {
	calculoArea() float64
}

func info (f figure) {
	value := f.calculoArea()
	fmt.Println("Area: ", value)
}

func main() {
	circle1 := circle {raio: 2}
	square1 := square {side: 4}

	info(circle1)
	info(square1)
}