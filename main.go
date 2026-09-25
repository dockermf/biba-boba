package main

import (
	_ "github.com/jackc/pgx/v5"
)

type Account struct {
	ID       uint64
	Login    string
	Password string
	Balance  float64
}

func main() {

}
