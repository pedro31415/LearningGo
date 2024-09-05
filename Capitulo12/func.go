package main

import "fmt"

//TERCEIRA AULA

type person struct {
	name string
	age int
}

// metoodo em Golang
func (p person) hello() {
	fmt.Println(p.name,", Ola")
}


// Estudo de interface e polimorfismo
type pessoa struct {
	name string
	sobreNome string
	age int
}

type dentista struct {
	pessoa pessoa
	denteArrancado int
	salario  int
}

type arquiteto struct {
	pessoa
	tipoDeConstrucao string
	tempoDeExperiencia int
}

func (x dentista) oiBomdia() {
	fmt.Println("Ola, eu sou ", x.pessoa.name, "e bom dia, eu arranquei ", x.denteArrancado)
}

func (x arquiteto) oiBomdia() {
	fmt.Println("Ola, eu sou ", x.pessoa.name, "e bom dia, eu tenho", x.tempoDeExperiencia, " anos de experiência")
}

type gente interface {
	oiBomdia()
}

func serHumano (g gente) {
	g.oiBomdia()
	switch g.(type) {
	case dentista: {
			fmt.Println("Eu ganho: ", g.(dentista).salario)
		}
	case arquiteto: {
			fmt.Println("Eu construo: ", g.(arquiteto).tipoDeConstrucao)
		}
	}
}

func main() {

	pedro := person {
		name: "Pedro",
		age: 21,
	}

	pedro.hello()

	mrDente := dentista {
		pessoa: pessoa{
			name: "Pedro",
			sobreNome: "Henrique",
			age: 21,
		},
		denteArrancado: 200,
		salario: 2000,
	}

	mrPredio := arquiteto {
		pessoa: pessoa{
			name: "Carlos",
			sobreNome: "Santos",
			age: 21,
		},
		tipoDeConstrucao: "Predios",
		tempoDeExperiencia: 20,
	}

	mrDente.oiBomdia()
	mrPredio.oiBomdia()

	serHumano(mrDente)
	fmt.Println()
	serHumano(mrPredio)

	fmt.Println(mrDente, mrPredio)
	// defer fmt.Println("Hello Wolrd")
	value := basica()
	fmt.Println(value)

	finalSum := adicao(12,6)
	fmt.Println(finalSum)

	somatorio, elementos := sum(30,30,30,30,30)
	fmt.Println(somatorio,elementos)

	slice1 := []int {
		90,60,70,39,34,70,10,
	}

	GaussSum(slice1)
}

// PRIMEIRA AULA
func basica() string {
	return "hoje é domingo"
}

func adicao(x int, y int) int {
	return x + y
}

func sum(x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}

// SEGUNDA AULA

func GaussSum (x[] int) {
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i]);
	}
}




