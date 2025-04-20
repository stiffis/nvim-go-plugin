package main

import (
	"log"
	"os"

	"github.com/neovim/go-client/nvim"
)

func main() {
	// Conectar al cliente Neovim
	vim, err := nvim.Dial(os.Getenv("NVIM_LISTEN_ADDRESS"))
	if err != nil {
		log.Fatalf("Error al conectar con Neovim: %v", err)
	}
	defer vim.Close()

	// Registrar un comando personalizado
	err = vim.RegisterHandler("HelloWorld", func(args []string) {
		// Manejar el error de vim.Command
		if err := vim.Command("echo '¡Hola desde Go!'"); err != nil {
			log.Printf("Error ejecutando comando: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Error al registrar el comando: %v", err)
	}

	// Manejar el error de vim.Serve
	if err := vim.Serve(); err != nil {
		log.Fatalf("Error al servir eventos de Neovim: %v", err)
	}
}
