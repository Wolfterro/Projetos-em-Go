package main

import (
	"fmt"
)

func main() {
	version := 1.0
	showMenu(version)
}

func showMenu(version float64) {
	var language int

	fmt.Println("========================")
	fmt.Printf("Hello World - Versão %.1f\n", version)
	fmt.Printf("========================\n\n")
	
	fmt.Println("Selecione o idioma:")
	fmt.Println("1 - Português")
	fmt.Println("2 - Inglês")
	fmt.Println("3 - Espanhol")
	fmt.Println("4 - Francês")
	fmt.Println("5 - Italiano")
	fmt.Println("6 - Alemão")
	fmt.Println("7 - Japonês")
	fmt.Println("8 - Chinês Tradicional")
	fmt.Println("9 - Russo")
	fmt.Println("0 - Sair do Programa")

	fmt.Printf("\nEscolha uma das opções acima: ")
	fmt.Scan(&language)
	fmt.Println("")

	printHello(language)
}

func printHello(language int) {
	PORTUGUES := "Olá, mundo!\n"
	INGLES := "Hello, world!\n"
	ESPANHOL := "¡Hola, mundo!\n"
	FRANCES := "Bonjour tout le monde!\n"
	ITALIANO := "Ciao, mondo!\n"
	ALEMAO := "Hallo, Welt!\n"
	JAPONES := "こんにちは、世界! (Kon'nichiwa, sekai!)\n"
	CHINESTRAD := "世界，你好！(Shìjiè, nǐ hǎo!)\n"
	RUSSO := "Привет, мир! (Privet, mir!)\n"

	switch(language) {
		case 1:
			fmt.Println(PORTUGUES)
		case 2:
			fmt.Println(INGLES)
		case 3:
			fmt.Println(ESPANHOL)
		case 4:
			fmt.Println(FRANCES)
		case 5:
			fmt.Println(ITALIANO)
		case 6:
			fmt.Println(ALEMAO)
		case 7:
			fmt.Println(JAPONES)
		case 8:
			fmt.Println(CHINESTRAD)
		case 9:
			fmt.Println(RUSSO)
		default:
			return
	}
}