# README
Este é um programa escrito em Go que busca a cotação do dólar em relação ao real brasileiro. O programa faz uma requisição HTTP para o servidor de moedas em http://localhost:8080/cotacao e extrai o valor da cotação em tempo real. Em seguida, o programa salva o valor da cotação em um arquivo de texto chamado cotacao.txt.

# Pré-requisitos
Este programa requer que o Go esteja instalado em sua máquina. Você pode baixar a versão mais recente do Go em https://golang.org/dl/.

# Como executar
1- Clone este repositório em sua máquina.

2- Abra o terminal e navegue até o diretório do repositório clonado.

3- Execute o seguinte comando para compilar o programa:

    go build

4- Execute o programa com o seguinte comando:

    ./currency-converter
    
5- não se esqueca de certificar que o /server/main.go esteja rodando.
