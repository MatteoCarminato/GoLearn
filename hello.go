package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("Hello: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := Hello("Matteo")
	// Se um erro for retornado, imprima-o no console e
	// sai do programa.
	if err != nil {
		log.Fatal(err)
	}

	// Se nenhum erro foi retornado, imprime a mensagem retornada
	// para o console.
	fmt.Println(message)
}
// Hello retorna a mensagem com o nome passado por parametro
func Hello(name string) (string, error){
	// Se nenhum nome foi fornecido, retorna um erro com uma mensagem.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Se um nome foi recebido, retorna um valor que incorpora o nome
	// em uma mensagem de saudação.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

