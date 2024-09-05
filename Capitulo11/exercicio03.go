package main

import "fmt"


type veiculo struct{
	portas int
	cor string
}

type caminhonete struct {
	veiculo veiculo
	quatroRodas bool
}


type sedan struct {
	veiculo veiculo
	modeloLuxo bool
}

func main() {
	caminhonete1 := caminhonete {
		veiculo: veiculo{
			portas: 6,
			cor: "Rosa",
		},
		quatroRodas: false,
	}

	sedan1 := sedan {
		veiculo: veiculo{
			portas: 4,
			cor: "Azul",
		},
		modeloLuxo: false,
	}

	fmt.Println(caminhonete1)
	fmt.Println(sedan1)
	fmt.Println("Portas da camionhe -> ", caminhonete1.veiculo.portas)
	fmt.Println("Cor do Sedan -> ", sedan1.veiculo.cor)

}