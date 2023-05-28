package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	dictionary := make(map[string]string)

	fmt.Println("Bem-vindo à ferramenta de criação de dicionário")
	fmt.Println("A ferramenta continuará pedindo para que você adicione palavras até que você digite NULL")

	for entry := 0; entry < 1; entry-- {
		var word string
		var meaning string

		fmt.Println("---------------------------------------------")
		fmt.Print("Palavra: ")
		fmt.Scanln(&word)

		if word == "NULL" {
			break
		}

		fmt.Print("Significado: ")
		fmt.Scanln(&meaning)

		dictionary[word] = meaning
	}

	clearScreen()

	for word, meaning := range dictionary {
		fmt.Println(word + ": " + meaning)
	}
}

func clearScreen() {
	var cmd *exec.Cmd

	cmd = exec.Command("clear")

	cmd.Stdout = os.Stdout
	cmd.Run()
}
