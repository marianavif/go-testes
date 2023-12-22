package ponteiros

import "testing"

func TestCarteira(t *testing.T) {
	confirmaSaldo := func(t *testing.T, carteira Carteira, esperado Bitcoin) {
		t.Helper()
		resultado := carteira.Saldo()
		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		esperado := Bitcoin(10)
		confirmaSaldo(t, carteira, esperado)

	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		carteira.Retirar(Bitcoin(10))
		esperado := Bitcoin(10)
		confirmaSaldo(t, carteira, esperado)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)

		if erro == nil {
			t.Error("Esperava um erroo mas nenhum ocorreu")
		}
	})
}
