package slices

func Soma(arr []int) int {
	var soma int
	for _, numero := range arr {
		soma += numero
	}
	return soma
}
