# Use a imagem oficial do Go como base
FROM golang:1.20-alpine

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o arquivo go.mod e go.sum para o contêiner
COPY go.mod .
COPY go.sum .

# Baixe e instale as dependências do Go
RUN go mod download

# Copie o diretório cmd para o contêiner
COPY cmd/ cmd/

# Copie o diretório internal para o contêiner
COPY internal/ internal/

# Compile o código fonte dentro do diretório cmd/seu-app
RUN go build -o /app/main ./cmd/main.go

# Expor a porta em que o servidor Go está escutando (se aplicável)
# EXPOSE 8080

# Comando para executar o seu aplicativo quando o contêiner for iniciado
CMD ["/app/main"]
