package main

import "fmt"

type Configuracoes struct {
	Porta   int    `json:"porta"`
	Host    string `json:"host"`
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
	Banco   string `json:"banco"`
}

func main() {

	configs := Configuracoes{
		Porta:   3306,
		Host:    "localhost",
		Usuario: "root",
		Senha:   "root",
		Banco:   "dev",
	}

	fmt.Println(configs)
}
