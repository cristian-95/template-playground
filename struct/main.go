package main

import "time"

type Voo struct {
	id       int
	origem   string
	destino  string
	partida  time.Time
	chegada  time.Time
	duracao  time.Time
	poltrona int
	classe   string
}

type User struct {
	id        int
	nome      string
	sobrenome string
	cpf       string
	idade     int
	premium   bool
	Voo
}

func main() {

}
