# Golang - API

API em go para gerenciamento de produtos, utilizando Docker e Postgres rodando em container

## Recursos

- Repositories 
- Usecases
- Controllers
  
## Rotas

- Lista produtos - http://localhost:8000/products [GET]
- Cadastra produto - http://localhost:8000/products [POST]
- Obtem um produto - http://localhost:8000/products/1 [GET]
- Atualiza produto - http://localhost:8000/products/1 [PUT]
- Exclui produto - http://localhost:8000/products/1 [DELETE]

## Subir app

Para subir a aplicação primeiro rode o comando para gerar a imagem "docker build -t go-api-vige ." e em seguida o comando "docker compose up -d" para subir o container com o app e o postgres.

```
$ docker build -t go-api-vige .
$ docker compose up -d
```

by [https://www.youtube.com/@GoLabTutoriais](https://www.youtube.com/@GoLabTutoriais)
Font: [https://www.youtube.com/watch?v=3p4mpId_ZU8&t=1875s] (Como criar uma REST API completa do zero com GO | Golang tutorial - iniciante)