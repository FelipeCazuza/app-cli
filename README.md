# 🚀 CLI de Busca DNS & Network

Uma aplicação de linha de comando (*CLI*) desenvolvida em **Go (Golang)** para realizar consultas de infraestrutura de rede e auditoria de registros DNS (IPs, Name Servers, TXT e CNAME).

## 📌 Funcionalidades

- **`ip`**: Retorna os endereços IPv4 e IPv6 associados a um host.
- **`servidores`**: Lista os Name Servers (NS) autoritativos do domínio.
- **`txt`**: Consulta registros TXT (SPF, verificações de domínio, etc.).
- **`cname`**: Exibe o registro CNAME (alias) de um host.

## 🛠️ Tecnologias e Ferramentas

- **Go 1.23**
- [urfave/cli](https://github.com/urfave/cli) - Framework para construção de CLIs em Go
- **Docker** (Multi-stage build com imagem `alpine`)
- **GitHub Actions** (Pipeline de CI para testes e build)

## 🧪 Executando os Testes Unitários

Para rodar a suíte de testes unitários:

```bash
go test ./app/... -v