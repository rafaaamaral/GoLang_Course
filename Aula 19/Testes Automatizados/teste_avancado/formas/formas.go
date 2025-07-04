package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func main() {
	r := Retangulo{Altura: 10, Largura: 5}
	c := Circulo{Raio: 5}

	EscreverArea(r)
	EscreverArea(c)
}
