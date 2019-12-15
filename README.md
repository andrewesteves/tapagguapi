# TaPaggu API

## Introduction

## Routes:

### Guard
/companies/{id} GET

/companies/{id} PUT

/companies/{id} DELETE

/companies POST

/companies GET

/categories/{id} GET

/categories/{id} PUT

/categories/{id} DELETE

/categories POST

/categories GET

/users/logout POST

/users/{id} GET

/users/{id} PUT

/users/{id} DELETE

/users GET

/receipts/query/{field} GET

/receipts/retrieve" GET

/receipts/{id} GET

/receipts/{id} PUT

/receipts/{id} DELETE

/receipts POST

/receipts GET

/receipts/{receiptID}/items/{id} GET

/receipts/{receiptID}/items/{id} PUT

/receipts/{receiptID}/items/{id} DELETE

/receipts/{receiptID}/items POST

/receipts/{receiptID}/items GET

### Unguard
/users/login POST

/users/recover POST

/users POST

## Create an user

### Request
Method: POST

Endpoint: /users

Body:

```json
{
  "name": "Andrew Luiz",
  "email": "andrewluiz@tapaggu.com",
  "password": "1234567"
}
```

### Response
```json
{
  "id": 1,
  "name": "Andrew Luiz",
  "email": "andrewluiz@tapaggu.com",
  "token": "JDJhJDEwJG[...]"
}
```

## Retrieve and store a receipt

### Request
Method: POST

Endpoint: /receipts/retrieve?url=URL

Authorization: Bearer Token

### Response
```json
{
  "receipt": {
    "id": 2,
    "category": {
      "id": 1,
      "title": "Geral",
      "icon": "all"
    },
    "company": {
      "id": 1,
      "cnpj": "05994622000368",
      "name": "JPLL COMERCIO E SERVICOS EIRELI",
      "title": "",
      "address": "RUA PROFESSOR RUI BATISTA, 120, B VIAGEM, Recife, PE, 51020160"
    },
    "title": "",
    "tax": 16.37,
    "extra": 9.1,
    "discount": 0,
    "total": 100.1,
    "items": [
      {
        "id": 7,
        "title": "CAPRI",
        "price": 23.9,
        "quantity": 1,
        "total": 23.9,
        "tax": 4.3,
        "measure": "UN",
        "created_at": "2019-12-11T20:38:08.699341Z",
        "updated_at": "2019-12-11T20:38:08.699341Z"
      },
      {
        "id": 5,
        "title": "COCA COLA",
        "price": 6.3,
        "quantity": 1,
        "total": 6.3,
        "tax": 1.13,
        "measure": "UN",
        "created_at": "2019-12-11T20:38:08.699318Z",
        "updated_at": "2019-12-11T20:38:08.699319Z"
      },
      {
        "id": 8,
        "title": "MONTPELLIER NUTELLA  S",
        "price": 29.9,
        "quantity": 1,
        "total": 29.9,
        "tax": 5.38,
        "measure": "UN",
        "created_at": "2019-12-11T20:38:08.699344Z",
        "updated_at": "2019-12-11T20:38:08.699344Z"
      },
      {
        "id": 6,
        "title": "PERPIGNAN",
        "price": 30.9,
        "quantity": 1,
        "total": 30.9,
        "tax": 5.56,
        "measure": "UN",
        "created_at": "2019-12-11T20:38:08.699338Z",
        "updated_at": "2019-12-11T20:38:08.699338Z"
      }
    ],
    "url": "http://nfce.sefaz.pe.gov.br/nfce-web/consultarNFCe?p=26191105994622000368650010000954761580180810",
    "accessKey": "26191105994622000368650010000954761580180810",
    "issuedAt": "2019-11-17T19:25:06Z",
    "createdAt": "2019-12-11T23:38:08.707119Z",
    "updatedAt": "2019-12-11T23:38:08.707119Z"
  }
}
```

## New receipt

### Request
Method: POST

Endpoint: /receipts

Authorization: Bearer Token


```json
{
  "category": {
    "title": "Transporte",
    "icon": "car"
  },
  "company": {
    "title": "Posto BR"
  },
  "title": "Viagem para arcoverde",
  "total": 200,
  "issuedAt": "2019-12-11T20:38:08.699341Z"
}
```

### Response
```json
{
  "receipt": {
    "id": 4,
    "category": {
      "id": 2,
      "title": "Transporte",
      "icon": "car"
    },
    "company": {
      "id": 3,
      "cnpj": "",
      "name": "",
      "title": "Posto BR",
      "address": ""
    },
    "title": "Viagem para Arcoverde",
    "tax": 0,
    "extra": 0,
    "discount": 0,
    "total": 200,
    "items": null,
    "url": "",
    "accessKey": "",
    "issuedAt": "2019-12-11T20:38:08.699341Z",
    "createdAt": "0001-01-01T00:00:00Z",
    "updatedAt": "0001-01-01T00:00:00Z"
  }
}
```

## New item from a receipt

### Request
Method: POST

Endpoint: /receipts/{receiptID}/items

Authorization: Bearer Token


```json
{
  "title": "Parabrisa",
  "price": 30,
  "quantity": 2,
  "total": 60,
  "measure": "UN"
}
```

### Response
```json
{
  "id": 10,
  "title": "Parabrisa",
  "price": 30,
  "quantity": 2,
  "total": 60,
  "tax": 0,
  "measure": "UN",
  "created_at": "0001-01-01T00:00:00Z",
  "updated_at": "0001-01-01T00:00:00Z"
}
```

## All receipts

### Request
Method: GET

Endpoint: /receipts

Params: month|year|category

Authorization: Bearer Token

### Response
```json
{
  "receipts": [
    {
      "id": 4,
      "category": {
        "id": 3,
        "title": "Transporte",
        "icon": "car",
        "total": 0
      },
      "company": {
        "id": 2,
        "cnpj": "",
        "name": "",
        "title": "Posto BR",
        "address": "  "
      },
      "title": "Viagem para arcoverde",
      "tax": 0,
      "extra": 0,
      "discount": 0,
      "total": 200,
      "items": [],
      "url": "",
      "accessKey": "",
      "issuedAt": "2019-12-12T20:38:08.699341Z",
      "createdAt": "2019-12-15T12:40:01.929391Z",
      "updatedAt": "2019-12-15T12:40:01.929391Z"
    },
    {
      "id": 2,
      "category": {
        "id": 2,
        "title": "Alimentação",
        "icon": "food",
        "total": 0
      },
      "company": {
        "id": 2,
        "cnpj": "",
        "name": "",
        "title": "Mc Donalds",
        "address": "  "
      },
      "title": "Viagem para arcoverde",
      "tax": 0,
      "extra": 0,
      "discount": 0,
      "total": 100,
      "items": [],
      "url": "",
      "accessKey": "",
      "issuedAt": "2019-12-11T20:38:08.699341Z",
      "createdAt": "2019-12-15T02:30:28.186413Z",
      "updatedAt": "2019-12-15T02:30:28.186413Z"
    },
    {
      "id": 1,
      "category": {
        "id": 2,
        "title": "Alimentação",
        "icon": "food",
        "total": 0
      },
      "company": {
        "id": 1,
        "cnpj": "",
        "name": "",
        "title": "Posto BR",
        "address": "  "
      },
      "title": "Viagem para arcoverde",
      "tax": 0,
      "extra": 0,
      "discount": 0,
      "total": 200,
      "items": [],
      "url": "",
      "accessKey": "",
      "issuedAt": "2019-12-11T20:38:08.699341Z",
      "createdAt": "2019-12-15T02:29:57.86471Z",
      "updatedAt": "2019-12-15T02:29:57.86471Z"
    }
  ],
  "categories": [
    {
      "id": 2,
      "title": "Alimentação",
      "icon": "food",
      "total": 300
    },
    {
      "id": 3,
      "title": "Transporte",
      "icon": "car",
      "total": 200
    }
  ],
  "current": "",
  "prev": "",
  "next": "",
  "total": ""
}
```

## New company

### Request
Method: POST

Endpoint: /companies

Authorization: Bearer Token


```json
{
  "cnpj": "34270953000140",
  "name": "A.L. Solutions",
  "title": "TaPaggu",
  "street": "RUA PROFESSOR RUI BATISTA",
  "number": "120",
  "district": "B VIAGEM",
  "city": "Recife",
  "state": "PE",
  "zipcode": "51020160"
}
```

### Response
```json
{
  "id": 2,
  "cnpj": "21270953000140",
  "name": "A.L. Solutions",
  "title": "TaPaggu",
  "street": "RUA PROFESSOR RUI BATISTA",
  "number": "120",
  "district": "B VIAGEM",
  "city": "Recife",
  "state": "PE",
  "zipcode": "51020160"
}
```

## New category

### Request
Method: POST

Endpoint: /categories

Authorization: Bearer Token


```json
{
  "title": "Transporte",
  "icon": "car"
}
```

### Response
```json
{
  "id": 3,
  "title": "Atividade",
  "icon": "bike"
}
```