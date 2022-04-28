package usecases

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func BuscarServidoresUsecase() cli.Command {
	command := cli.Command{
		Name:  "servidores",
		Usage: "Busca nome dos servidores na internet",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Value: "devbook.com.br",
			},
		},
		Action: buscarServidores,
	}

	return command
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for idx, servidor := range servidores {
		fmt.Printf("Host: %s|| SERVIDOR%v: %v\n", host, idx+1, servidor.Host)
	}
}
