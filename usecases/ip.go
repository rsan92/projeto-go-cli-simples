package usecases

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func BuscarIpsUsecase() cli.Command {
	command := cli.Command{
		Name:  "ip",
		Usage: "Busca IPs de endereços na internet",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Value: "devbook.com.br",
			},
		},
		Action: buscarIps,
	}

	return command
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for idx, ip := range ips {
		fmt.Printf("Host: %s|| IP%v: %v\n", host, idx+1, ip)
	}
}
