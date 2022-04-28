package app

import (
	"go-cli-test/usecases"

	"github.com/urfave/cli"
)

// Gerar: Retorna app de linha de comando, pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()

	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs de servidores da internet"

	app.Commands = []cli.Command{}

	app.Commands = append(app.Commands, usecases.BuscarIpsUsecase())

	app.Commands = append(app.Commands, usecases.BuscarServidoresUsecase())

	return app

}
