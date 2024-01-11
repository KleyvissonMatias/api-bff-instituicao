# API BFF Instituição

Bem-vindo ao repositório da API BFF Instituição, um projeto escrito em Go, utilizando a versão mais recente, e MongoDB como banco de dados. Esta API serve como um Backend For Frontend (BFF) para gerenciar operações CRUD (Create, Read, Update, Delete) relacionadas a instituições de ensino.

## Visão Geral

A API BFF Instituição foi desenvolvida para oferecer uma camada de serviço eficiente e especializada para operações relacionadas a instituições de ensino. Ela utiliza Go para alta performance e MongoDB para armazenamento de dados.

## Tecnologias Utilizadas

- Go (versão mais recente)
- MongoDB

## Configuração do Ambiente

Certifique-se de ter o Go e o MongoDB instalados em seu ambiente de desenvolvimento. Além disso, você precisará configurar as credenciais do banco de dados no arquivo `.toml`.

```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=nome_do_banco
```

## Como Executar

1. Clone o repositório: `git clone https://github.com/seu-usuario/api-bff-instituicao.git`
2. Navegue até o diretório do projeto: `cd api-bff-instituicao`
3. Execute a aplicação: `go run main.go`

A API estará disponível em `http://localhost:8080`.

## Endpoints Disponíveis

- `GET /api/v1/instituicoes`: Retorna todas as instituições cadastradas.
- `GET /api/v1/instituicoes/{id}`: Retorna os detalhes de uma instituição específica.
- `POST /api/v1/instituicoes`: Cria uma nova instituição.
- `PUT /api/v1/instituicoes/{id}`: Atualiza os dados de uma instituição existente.
- `DELETE /api/v1/instituicoes/{id}`: Exclui uma instituição pelo ID.

## Estrutura do Projeto

A estrutura do projeto segue um padrão organizacional claro para facilitar a manutenção e expansão:

- `main.go`: Ponto de entrada da aplicação.
- `api/`: Arquivo RAML.
- `models/`: Definição de modelos de dados.
- `services/`: Lógica de negócios e interação com o banco de dados.
- `repository/`: Camada de acesso ao banco de dados.
