package maps

import "errors"

type Dicionario map[string]string

var ErrNaoEncontrado = errors.New("não foi possível encontrar a palavra que você procura")

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) {
	d[palavra] = definicao
}
