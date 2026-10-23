package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar creates and returns a new CLI application instance.
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPS e Nomes de servidores na internet "
	app.Version = "1.0.0"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS da internet",
			Flags:  flags,
			Action: buscarIps,
		},

		{
			Name:   "servidores",
			Usage:  "Busca nomes de servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},

		{
			Name:   "txt",
			Usage:  "Busca registros TXT na internet",
			Flags:  flags,
			Action: buscarTxt,
		},

		{
			Name:   "cname",
			Usage:  "Busca registro CNAME na internet",
			Flags:  flags,
			Action: buscarCname,
		},
	}
	return app
}

func buscarIps(c *cli.Context) error {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("Erro ao buscar IPS para o host %s: %v", host, err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}
func buscarServidores(c *cli.Context) error {
	host := c.String("host")

	servidores, err := net.LookupNS(host) // name server
	if err != nil {
		fmt.Printf("Não foi possível encontrar registros NS para '%s'. Dica: verifique se informou o domínio raiz (ex: dominio.com.br) em vez de um subdomínio.\n", host)

	}
	for _, server := range servidores {
		fmt.Println(server.Host)
	}
	return nil
}

func buscarTxt(c *cli.Context) error {
	host := c.String("host")

	txts, err := net.LookupTXT(host)
	if err != nil {
		log.Fatalf("Erro ao buscar registros TXT para o host %s: %v", host, err)
	}
	for _, txt := range txts {
		fmt.Println(txt)
	}
	return nil
}

func buscarCname(c *cli.Context) error {
	host := c.String("host")

	cname, err := net.LookupCNAME(host)
	if err != nil {
		log.Fatalf("Erro ao buscar registro CNAME para o host %s: %v", host, err)
	}
	fmt.Println(cname)
	return nil
}
