package main

import "fmt"

func main() {
	nome := "Paulo"
	versao := 1.2
	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("O comando escolhindo foi", comando)
}
