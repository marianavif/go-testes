package maps

import (
	"testing"
)

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é apenas um teste"

		comparaStrings(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, err := dicionario.Busca("desconhecida")

		if err == nil {
			t.Fatal("é esperado que um erro seja obtido")
		}
	})
}

func TestAdiciona(t *testing.T) {
	dicionario := Dicionario{}
	palavra := "teste"
	definicao := "isso é apenas um teste"
	dicionario.Adiciona(palavra, definicao)

	comparaDefinicao(t, dicionario, palavra, definicao)

}

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s' ", resultado, esperado, "teste")
	}
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s' ", resultado, esperado)
	}
}

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()

	resultado, err := dicionario.Busca(palavra)
	if err != nil {
		t.Fatal("não foi possível encontrar a palavra adicionada:", err)
	}
	if definicao != resultado {
		t.Errorf("resultado '%s', esperado '%s' ", resultado, definicao)
	}
}
