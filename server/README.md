# Server cotação de moedas com SQLite

Este é server API para cotação de moedas com armazenamento em banco de dados SQLite.

## Instalação

Para executar esse exemplo, é necessário ter o SQLite instalado em sua máquina. Caso não tenha, você pode fazer o download através do [site oficial](https://www.sqlite.org/download.html) ou através do gerenciador de pacotes da sua distribuição.

### Instalando SQLite no Ubuntu

Para instalar o SQLite no Ubuntu, você pode executar o seguinte comando:

    sudo apt-get install sqlite3

## Executando o projeto

Para criar a tabela currencies no banco de dados SQLite, você pode executar o seguinte comando no terminal na raiz do projeto:


    $ sqlite3 currency.db < migration.sql
Este comando executará o arquivo migration.sql, que contém a query SQL para criar a tabela currencies.

Antes de executar o código main.go, é necessário instalar as dependências. Para isso, execute o seguinte comando na raiz do projeto:
    
    $ go mod tidy

Em seguida, execute o seguinte comando para compilar e executar o código:

    $ go run main.go

O server estará disponível em:

     GET - http://localhost:8080/cotacao