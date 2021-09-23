package main

import (
	"fmt"
	"log"
)

type Falante interface {
	Falar(texto string) string
}

type Pessoa struct {
	Nome string
}

func (p Pessoa) Falar(texto string) string {
	return fmt.Sprintf("olá, me chamo %s. %s", p.Nome, texto)
}

type Papagaio struct {
	Nome string
}

func (p Papagaio) Falar(texto string) string {
	return fmt.Sprintf("olá, me chamo %s. %s", p.Nome, texto)
}

func falar(f Falante) {
	log.Println(f.Falar("teste"))
}

func mainTest() {
	papagaio := Papagaio{
		Nome: "Loro",
	}

	pessoa := Pessoa{
		Nome: "Anderson",
	}

	falar(papagaio)
	falar(pessoa)
}
