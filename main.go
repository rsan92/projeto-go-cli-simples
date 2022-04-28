package main

import (
	"fmt"
	"go-cli-test/app"
	"log"
	"os"
)

func main() {
	fmt.Println("-- PGM Iniciado --")
	app := app.Gerar()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
