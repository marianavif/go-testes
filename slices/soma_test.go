package slices

import "testing"

func TestSoma(t *testing.T) {

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		soma := Soma(arr)
		esperado := 15

		if soma != esperado {
			t.Errorf("esperado '%d' mas obteve '%d' dado %v ", esperado, soma, arr)
		}
	})
}
