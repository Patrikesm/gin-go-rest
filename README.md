## Tecnologias utilizadas

- GIN como principal criado da API, funciona como um express
- GORM como ORM, funciona como um sequelize
- Docker para o postgres
- Go valdator para validações de campos e outros - gopkg.in/validator.v2
- Testify para auxiliar nos testes unitários - github.com/stretchr/testify

## Principais funções

- Criar uma API Rest, que realiza um CRUD com banco postgres

## Principais utilizações do GIN

- Criar parte de rotas
- Iniciar servidor
- Definir config padrão para API

## Principais utilizações GORM

- Criar models com automigração
- Realizar conexão com banco de dados
- Realizar todas as inserções, deleções e alterações no banco de dados

## Principais utilizações docker

- subir um pgAdmin
- subir uma base de dados postgres


## Como rodar os testes em golang

- abra o terminal e digite os seguintes comandos:

GO TEST - quando quiser testar todos
GOR TEST -RUN {NomeDoTeste} - quando quiser testar um em especifico