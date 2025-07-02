package formas

import "testing"

func TestAre(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		r := Retangulo{Altura: 10, Largura: 5}
		if r.area() != 50 {
			t.Errorf("Área do retângulo incorreta, esperado 50, obtido %f", r.area())
		}
	})
	t.Run("Circulo", func(t *testing.T) {
		c := Circulo{Raio: 5}
		if c.area() != 78.53981633974483 {
			t.Errorf("Área do círculo incorreta, esperado 78.54, obtido %f", c.area())
		}
	})
}
