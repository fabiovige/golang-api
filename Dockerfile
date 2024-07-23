# Usar a imagem base do Golang
FROM golang:latest

# Configurar o diretório de trabalho
WORKDIR /go/src/app

# Copiar apenas o go.mod e go.sum para utilizar cache de dependências
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o restante do código
COPY . .

# Expor a porta que a aplicação vai rodar
EXPOSE 8000

# Construir a aplicação Go
RUN go build -o main ./cmd/main.go

# Executar a aplicação
CMD ["./main"]
