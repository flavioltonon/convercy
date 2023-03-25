# Convercy

Uma API de conversão de moedas escrita em Go.

## Pré-requisitos

- golang@v1.18 ou mais recente
- docker-compose@v2.15.1

## Funcionalidades

### Gerenciamento de moedas registradas

A aplicação permite o controle de moedas registradas pelo usuário. Moedas registradas representam quais resultados de conversão de moedas deverão ser entregues pela funcionalidade de conversão.

### Convertendo uma quantidade de moeda

Com o Convercy, você pode fazer a conversão de quantidades de uma moeda específica para todos os tipos de moeda que você tiver registrado.

## Rodando o projeto

### Sistemas operacionais baseados em Unix

> make start

### Outros sistemas operacionais

> docker-compose up

## Documentação

### API

A especificação completa da API pode ser encontrada em docs/openapi.json. Esta especificação pode ser facilmente importada em ferramentas como o Postman ou o Insomnia.

### Diagramas de sequência

Diagramas de sequência dos fluxos da aplicação podem ser encontrados no diretório docs/websequencediagrams.

## Rodando testes

### Sistemas operacionais baseados em Unix

> make tests

### Outros sistemas operacionais

> go test -cover ./application/services/... ./domain/services/...

## Roadmap

- [ ] Adicionar camada de cache para reduzir a quantidade de chamadas para a API de taxas de câmbio da OpenExchangeRates (https://openexchangerates.org)
- [ ] Adicionar camada de autenticação para controlar o gerenciamento de moedas registradas por client
- [x] Adicionar pipeline de CI para testar alterações feitas na branch main e publicar imagens Docker quando uma release for criada