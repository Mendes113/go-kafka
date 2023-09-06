# Use a imagem base oficial do Go
FROM golang:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o código-fonte do seu aplicativo para o contêiner
COPY ./cmd ./cmd   
COPY . .

# Compile o aplicativo a partir do diretório 'cmd'
RUN go build -o main ./cmd

# Especifique o comando de inicialização
CMD ["./main"]
