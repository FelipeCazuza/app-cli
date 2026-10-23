package app

import (
	"flag"
	"testing"

	"github.com/urfave/cli"
)

// TestGerar verifica se a aplicação CLI é criada com as configurações corretas.
func TestGerar(t *testing.T) {
	app := Gerar()

	if app.Name != "Aplicação de Linha de Comando" {
		t.Errorf("Nome da aplicação incorreto. Esperado: 'Aplicação de Linha de Comando', Obtido: '%s'", app.Name)
	}

	if app.Version != "1.0.0" {
		t.Errorf("Versão da aplicação incorreta. Esperado: '1.0.0', Obtido: '%s'", app.Version)
	}

	if len(app.Commands) != 4 {
		t.Errorf("Número de comandos incorreto. Esperado: 4, Obtido: %d", len(app.Commands))
	}
}

// TestBuscarIps testa se a função de busca de IPs executa sem erros para um host válido
func TestBuscarIps(t *testing.T) {
	set := flag.NewFlagSet("test", flag.ContinueOnError)
	set.String("host", "google.com", "host para teste")
	ctx := cli.NewContext(nil, set, nil)

	err := buscarIps(ctx)
	if err != nil {
		t.Errorf("Esperado nenhum erro ao buscar IPs, mas obteve: %v", err)
	}
}

// TestBuscarServidores testa se a função de busca de NS executa sem erros para um domínio válido
func TestBuscarServidores(t *testing.T) {
	set := flag.NewFlagSet("test", flag.ContinueOnError)
	set.String("host", "google.com", "host para teste")
	ctx := cli.NewContext(nil, set, nil)

	err := buscarServidores(ctx)
	if err != nil {
		t.Errorf("Esperado nenhum erro ao buscar servidores NS, mas obteve: %v", err)
	}
}

func TestBuscarTxt(t *testing.T) {
	set := flag.NewFlagSet("test", flag.ContinueOnError)
	set.String("host", "google.com", "host para teste")
	ctx := cli.NewContext(nil, set, nil)

	err := buscarTxt(ctx)
	if err != nil {
		t.Errorf("Esperado nenhum erro ao buscar registros TXT, mas obteve: %v", err)
	}
}

func TestBuscarCname(t *testing.T) {
	set := flag.NewFlagSet("test", flag.ContinueOnError)
	set.String("host", "www.google.com", "host para teste")
	ctx := cli.NewContext(nil, set, nil)

	err := buscarCname(ctx)
	if err != nil {
		t.Errorf("Esperado nenhum erro ao buscar registros CNAME, mas obteve: %v", err)
	}
}
