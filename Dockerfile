# 01- Build da aplicação
FROM golang:1.23-alpine AS build

WORKDIR /app

#02- Copia os arquivos go.mod e go.sum e baixa as dependências
COPY go.mod go.sum ./
RUN go mod download

# 03- Copia o restante do código da aplicação
COPY . .

# 04- Compila um binário estático e otimizado para Linux
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app-cli main.go

# 05- Cria uma imagem mínima para rodar a aplicação
FROM alpine:latest

# 06- Copia o binário compilado da etapa anterior
COPY --from=build /app/app-cli .    

#Define o ponto de entrada da aplicação
ENTRYPOINT ["./app-cli"]