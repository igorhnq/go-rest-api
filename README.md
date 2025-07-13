# Go REST API

API REST simples em Go para gerenciamento de produtos, desenvolvido como projeto de estudo em Go com Clean Architecture.

## 📋 Descrição

Esta é uma API REST que permite gerenciar produtos através de operações CRUD básicas. O projeto utiliza Go 1.24.5 com Gin como framework web e PostgreSQL como banco de dados.

## 🏗️ Arquitetura

O projeto segue os princípios da Clean Architecture, organizado nas seguintes camadas:

- **Controllers**: Responsáveis por receber requisições HTTP e retornar respostas
- **Use Cases**: Contêm a lógica de negócio da aplicação
- **Repositories**: Gerenciam o acesso aos dados
- **Models**: Definem as estruturas de dados
- **Database**: Configuração e conexão com o banco de dados

## 🚀 Tecnologias Utilizadas

- **Go 1.24.5**: Linguagem principal
- **Gin**: Framework web para criação da API REST
- **PostgreSQL**: Banco de dados relacional
- **Docker**: Containerização da aplicação
- **Docker Compose**: Orquestração de containers

## 📁 Estrutura do Projeto

```
go-rest-api/
├── cmd/
│   └── main.go                 # Ponto de entrada da aplicação
├── controllers/
│   └── product_controller.go   # Controladores HTTP
├── db/
│   └── conn.go                 # Configuração do banco de dados
├── models/
│   ├── product.go              # Modelo de produto
│   └── response.go             # Modelo de resposta
├── repositories/
│   └── product_repository.go   # Acesso aos dados
├── usecases/
│   └── product_usecase.go      # Lógica de negócio
├── docker-compose.yml          # Configuração do Docker Compose
├── Dockerfile                  # Configuração do Docker
├── go.mod                      # Dependências do Go
└── go.sum                      # Checksums das dependências
```

## 📋 Pré-requisitos

- Go 1.24.5 ou superior
- Docker e Docker Compose
- PostgreSQL (opcional, se não usar Docker)

## 📋 Como Executar

### Com Docker (Recomendado)
```bash
docker-compose up --build
```

### Localmente
```bash
go mod tidy
go run cmd/main.go
```

A API estará disponível em `http://localhost:8080`

## 📋 Endpoints

- **GET** `/ping` - Health check
- **GET** `/products` - Lista todos os produtos
- **POST** `/products` - Cria um novo produto
- **GET** `/products/:id` - Busca produto por ID

## 📝 Exemplo de Uso

### Criar Produto
```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"nmproduct": "Produto Teste", "vlprice": 29.99}'
```

### Listar Produtos
```bash
curl -X GET http://localhost:8080/products
```

## 📊 Modelo de Dados

### Produto
```json
{
  "cdproduct": 1,
  "nmproduct": "Nome do Produto",
  "vlprice": 29.99
}
```

## 🔧 Variáveis de Ambiente

As configurações do banco de dados estão definidas no arquivo `db/conn.go`:

- `host`: go_db (Docker) ou localhost (local)
- `port`: 5432
- `user`: postgres
- `password`: 1234
- `dbname`: postgres

## 🐳 Docker

### Build da Imagem
```bash
docker build -t go-rest-api .
```

### Executar Container
```bash
docker run -p 8080:8080 go-rest-api
```
## 📦 Dependências

Principais dependências do projeto:

- `github.com/gin-gonic/gin`: Framework web
- `github.com/lib/pq`: Driver PostgreSQL

## 🔗 Links Úteis

- [Documentação do Go](https://golang.org/doc/)
- [Documentação do Gin](https://gin-gonic.com/docs/)
- [Documentação do PostgreSQL](https://www.postgresql.org/docs/)