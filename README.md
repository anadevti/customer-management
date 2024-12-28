## Uma API CRUD simples para gerenciamento de clientes, desenvolvida em Go. A API permite criar, ler, atualizar e excluir clientes, e pode ser testada com ferramentas como [Postman](https://www.postman.com/).

# Tecnologias Utilizadas

- **Go**: Linguagem de programação principal.
- **GORM**: ORM para integração com o banco de dados.
- **Banco de Dados**: PostgreSQL.
- **Docker** (opcional): Para containerização do banco.

# Endpoints da API
### Listar todos os clientes
GET /clients
Retorna uma lista de todos os clientes.

### Listar um cliente específico
GET /clients/{id}
Parâmetros: id (inteiro).
Retorna os dados do cliente especificado.

### Criar um novo cliente
POST /clients
Corpo (JSON)

### Atualizar um cliente
PUT /clients/{id}
Parâmetros: id (inteiro).
Corpo (JSON)

### Excluir um cliente
DELETE /clients/{id}
Parâmetros: id (inteiro).
Retorna uma mensagem de sucesso ou erro.

# Como Executar

- Clone o repositório
- Configure as variáveis de ambiente do banco:
```bash
export DB_HOST="localhost"
export DB_USER="root"
export DB_PASS="password"
export DB_NAME="clientes"
```
- Instale as dependências:
```bash
go mod tidy
```
- Execute a aplicação:
```bash
go run main.go
```
### A API estará disponível em http://localhost:8080
