package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// para converter para json os campos da struct necessita ficar em maiusculo
type pessoa struct {
	Nome string
	Sobrenome string
	Profissão string
	Contabilidade float64
}

// Unmarshal
type informacoes struct {
	Nome string `json:"Nome"`
	Sobrenome string `json:"Sobrenome"`
	Idade string `json:"Idade"`
	Profissão string `json:"Profissão"`
	ContaBancaria string `json:"ContaBancaria"`
}

func main(){
	james := pessoa {"James","Bond","Agente Secreto", 70000}

	sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Profissão":"Agente Secreto","Contabilidade":70000}`)

	var jamesbond informacoes
	err := json.Unmarshal(sb, &jamesbond)
	if err != nil {
		fmt.Println("Erro: ", err)
	}

	fmt.Println(jamesbond)
	fmt.Println(jamesbond.Nome)

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesbond)

	j, err := json.Marshal(james)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j))
}