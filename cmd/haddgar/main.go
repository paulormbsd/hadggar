package main

import (
	"log"

	"github.com/paulormbsd/hadggar/views"
)

// Função main para testar, cria uma mensagem de iniciar o serviço Zeus Hammer e chama a função de rota no final
func main() {
	log.Println("Starting Haddgar API")
	views.HandleRequest()
}
