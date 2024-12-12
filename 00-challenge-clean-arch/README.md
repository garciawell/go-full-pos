# Instruções para Rodar Desafio. 



## 1 - passo

Rodar o comando do docker para subir o container do mysql, banco de dados orders e rabbitMQ

```
    docker-compose  up -d
```

<br/>


## 2 - Rodar migration

Rodar as migrations para criar a tabela no DB

```
    make migrateUp
```

<br/>

## 3 - Rodar o projeto 

```
    make run
```

<br/>

## 4 - REST API 

Acessar a pasta api e executar o arquivo `order.http` para criar uma ordem e listar.

<br/>


## 5 - GRPC

Executar o comando para dar um call e utilizando as funções `CreateOrder` e `ListOrders`
```
    make evans
```


<br/>

## 6 - GraphQL

Acessar a URL http://localhost:8080 e executar os seguintes comandos para criar uma mutation e query

```graphql
mutation createOrder {
  createOrder(input: {id: "teste222", Price: 12, Tax: 12}) {
    id
    Price
    Tax
    FinalPrice
  }
}

query listOrder {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```