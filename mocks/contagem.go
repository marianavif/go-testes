package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// constants

const (
	ultimaPalavra  = "Vai!"
	inicioContagem = 3
	escrita        = "escrita"
	pausa          = "pausa"
)

// types

type Sleeper interface {
	Pausa()
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

// TempoSpy methods

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

// SleeperConfiguravel methods

func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

// SpyContagemOperacoes methods

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

// functions

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}
	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

// main function

func main() {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}
