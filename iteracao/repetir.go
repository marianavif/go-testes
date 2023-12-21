package iteracao

func Repetir(caractere string, quantidadeRepeticoes int) (repeticoes string) {
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}
	return
}
