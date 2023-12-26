package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	ultimaPalavra  = "Vai!"
	inicioContagem = 3
)

type Sleeper interface {
	Sleep()
}

type SleeperSpy struct {
	Chamadas int
}

type SleeperPadrao struct{}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}
	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}
